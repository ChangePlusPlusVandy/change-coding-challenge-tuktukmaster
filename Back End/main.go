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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnTweetFromHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: grab tweet page")

	enableCors(&w)

	vars := mux.Vars(r)
	key := vars["handle"]

	fmt.Println("Handle: " + key)

	tweets := tweets.GrabTweets(key)

	fmt.Fprintf(w, "[")
	for i := 0; i < len(tweets); i++ {
		printableTweet, _ := json.Marshal(tweets[i])
		fmt.Fprintf(w, string(printableTweet))
		if i != len(tweets)-1 {
			fmt.Fprintf(w, ",")
		}
	}
	fmt.Fprintf(w, "]")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleRequests(port string, bearer string) {
	fmt.Println("Hosting backend on port " + string(port))

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/get-tweets/{handle}", returnTweetFromHandle).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func main() {
	confFile, _ := ioutil.ReadFile("./conf.json")
	secretsFile, _ := ioutil.ReadFile("./secrets.json")

	conf := confJSON{}
	secrets := secretsJSON{}

	_ = json.Unmarshal(confFile, &conf)
	_ = json.Unmarshal(secretsFile, &secrets)

	handleRequests(conf.Port, secrets.Bearer)
}
