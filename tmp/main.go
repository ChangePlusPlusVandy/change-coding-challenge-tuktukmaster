package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"regexp"
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

func filter(tweetArr []TweetData, test func(TweetData) bool) (ret []TweetData) {
	for _, tweet := range tweetArr {
		if test(tweet) {
			ret = append(ret, tweet)
		}
	}
	return
}

func DoesNotContainLinkOrTag(tweet TweetData) bool {
	tag, _ := regexp.Compile("@\\S*")
	link, _ := regexp.Compile("https:\\/\\/\\S*")
	if tag.MatchString(tweet.Text) {
		//fmt.Println("TAG: " + tweet.Text)
		return false
	}
	if link.MatchString(tweet.Text) {
		//fmt.Println("LINK: " + tweet.Text)
		return false
	}
	return true
}

func grabTweets(name string) {
	var offset string = ""
	var allTweets = make([]TweetData, 3200)
	for i := 0; i < 3200; i += 200 {
		var tempTweets []TweetData
		var buildTweets []byte

		// Per twitter api, you can only grab a max of 200
		buildTweets = <-grabTwoHundredTweets(name, offset)

		// So load these 200 into the temp tweets array
		_ = json.Unmarshal([]byte(buildTweets), &tempTweets)

		// And copy them into the all tweets array that will be returned
		copy(allTweets[i:i+200], tempTweets[:])
	}

	// Now you need to filter all of the tweets to not include retweets and images
	// This could have mostly been done in the api call but that would vastly complicate
	// creating a nice array so I decided to filter it here
	filteredTweets := filter(allTweets, DoesNotContainLinkOrTag)

	for i := 0; i < len(filteredTweets); i++ {
		//fmt.Println(filteredTweets[i].Text)
	}
	fmt.Println(len(filteredTweets))
}

func main() {
	grabTweets("elonmusk")
	grabTweets("kanyewest")
	grabTweets("realDonaldTrump")
}
