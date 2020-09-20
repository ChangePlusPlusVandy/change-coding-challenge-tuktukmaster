package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
)

// UserData info
type UserData struct {
	ScreenName string `json:screen_name`
	Name       string `json:name`
}

type TweetData struct {
	ID_STR string   `json:id_str`
	Text   string   `json:text`
	User   UserData `json:user`
}

/*{
"created_at":"Mon Aug 24 22:32:59 +0000 2020",
"id":1298025540742524928,
"id_str":"1298025540742524928",
"text":"@Erdayastronaut Orbital launch mount",
"truncated":false,
"entities":{
	"hashtags":[],
	"symbols":[],
	"user_mentions":[
		{"screen_name":"Erdayastronaut",
		"name":"Everyday Astronaut",
		"id":3167257102,
		"id_str":"3167257102",
		"indices":[0,15]}
		],
		"urls":[]
	},
	"source":"\u003ca href=\"http:\/\/twitter.com\/download\/iphone\" rel=\"nofollow\"\u003eTwitter for iPhone\u003c\/a\u003e",
	"in_reply_to_status_id":1298022728944029698,
	"in_reply_to_status_id_str":"1298022728944029698",
	"in_reply_to_user_id":3167257102,
	"in_reply_to_user_id_str":"3167257102",
	"in_reply_to_screen_name":"Erdayastronaut",
	"user":{
		"id":44196397,
		"id_str":"44196397",
		"name":"Elon Musk",
		"screen_name":"elonmusk",
		"location":"",
		"description":"",
		"url":null,
		"entities":
		{"description":{"urls":[]}},
		"protected":false,
		"followers_count":38788634,
		"friends_count":97,
		"listed_count":55991,
		"created_at":"Tue Jun 02 20:12:29 +0000 2009",
		"favourites_count":6658,
		"utc_offset":null,
		"time_zone":null,
		"geo_enabled":false,
		"verified":true,
		"statuses_count":12319,
		"lang":null,"contributors_enabled":false,"is_translator":false,"is_translation_enabled":false,"profile_background_color":"C0DEED","profile_background_image_url":"http:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png","profile_background_image_url_https":"https:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png","profile_background_tile":false,"profile_image_url":"http:\/\/pbs.twimg.com\/profile_images\/1295975423654977537\/dHw9JcrK_normal.jpg","profile_image_url_https":"https:\/\/pbs.twimg.com\/profile_images\/1295975423654977537\/dHw9JcrK_normal.jpg","profile_banner_url":"https:\/\/pbs.twimg.com\/profile_banners\/44196397\/1576183471","profile_link_color":"0084B4","profile_sidebar_border_color":"C0DEED","profile_sidebar_fill_color":"DDEEF6","profile_text_color":"333333","profile_use_background_image":true,"has_extended_profile":true,"default_profile":false,"default_profile_image":false,"following":null,"follow_request_sent":null,"notifications":null,"translator_type":"none"},
		"geo":null,
		"coordinates":null,
		"place":null,
		"contributors":null,
		"is_quote_status":false,
		"retweet_count":293,
		"favorite_count":6192,
		"favorited":false,
		"retweeted":false,
		"lang":"en"
	}
*/
func grabTwoHundredTweets(name string, offsetStr string) <-chan []byte {
	r := make(chan []byte)

	go func() {
		defer close(r)
		url := "https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=" + name + "&count=200"

		if offsetStr != "" {
			// max_id is inclusive so you need to add 1 to grab 200 new tweets
			offsetID := new(big.Int)
			offsetID.SetString(offsetStr, 10)
			offsetID.Add(offsetID, big.NewInt(1))
			url += "&max_id=" + offsetID.String()
		}

		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + "AAAAAAAAAAAAAAAAAAAAADlOHwEAAAAAr9Ncm6oJ8hhQ5U18GafR9ORycH4%3DDfgEkAO74v0YRMLdEbm5LkmhADayzqN26PkUTcf70A5A1v5D56"

		// Create a new request using http
		req, err := http.NewRequest("GET", url, nil)

		// add authorization header to the req
		req.Header.Add("Authorization", bearer)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		}
		//log.Println(resp.Body)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("Error while reading body")
		}

		r <- []byte(body)
	}()

	return r
}

func main() {

	var tweets []TweetData
	var buildTweets []byte

	buildTweets = <-grabTwoHundredTweets("elonmusk", "")
	for i := 200; i < 3200; i += 200 {
		var tempTweets []TweetData
		_ = json.Unmarshal([]byte(buildTweets), &tempTweets)
		//fmt.Println(tempTweets[0].ID_STR)
		lastId := tempTweets[i-1].ID_STR
		buildTweets = append(buildTweets, <-grabTwoHundredTweets("elonmusk", lastId)...)
	}

	_ = json.Unmarshal([]byte(buildTweets), &tweets)
	fmt.Println(len(tweets))
	for i := 0; i < len(tweets); i++ {
		//fmt.Println(tweets[i].ID_STR)
	}
}
