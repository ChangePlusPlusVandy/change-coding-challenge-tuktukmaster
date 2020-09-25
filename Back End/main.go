package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./tweets"

	"github.com/gorilla/mux"
)

type confJSON struct {
	Port string `json:port`
}

type secretsJSON struct {
	Bearer string `string:bearer`
}

// home page largely for debugging but also doubles as a friendly message :)
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// page visited to actually retrieve
func returnTweetFromHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: grab tweet page")

	// enable CORS to allow browser to make call to API
	enableCors(&w)

	// grab the desired twitter handle from the request
	vars := mux.Vars(r)
	key := vars["handle"]

	// print out a helpful message to the console
	fmt.Println("Handle: " + key)

	// grab all the tweets from the tweets package
	tweets := tweets.GrabTweets(key)

	// Put it in an array format
	fmt.Fprintf(w, "[")
	for i := 0; i < len(tweets); i++ {
		// Marshall the TweetData so it can be printed out
		printableTweet, _ := json.Marshal(tweets[i])
		fmt.Fprintf(w, string(printableTweet))

		// Add commas on everything except for the last one bc it is an array
		if i != len(tweets)-1 {
			fmt.Fprintf(w, ",")
		}
	}
	fmt.Fprintf(w, "]")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleRequests(port string) {
	fmt.Println("Hosting backend on port " + string(port))

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/get-tweets/{handle}", returnTweetFromHandle).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func main() {
	confFile, _ := ioutil.ReadFile("./conf.json")

	conf := confJSON{}

	_ = json.Unmarshal(confFile, &conf)

	handleRequests(conf.Port)
}
