package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"unsafe"
)

type slack struct {
	Text       string `json:"text"`
	Username   string `json:"username"`
	Icon_emoji string `json:"icon_emoji"`
	Channel    string `json:"channel"`
}

func exists(f string) bool {
	_, err := os.Stat(f)
	return err == nil
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	slack_username := getenv("SLACK_POST_USERNAME", hostname)
	slack_icon := getenv("SLACK_POST_ICON", ":robot_face:")
	slack_post_channel := getenv("SLACK_POST_CHANNEL", "#bots")

	webhook_url := os.Getenv("SLACK_POST_WEBHOOK_URL")

	if len(webhook_url) == 0 {
		log.Fatal("Not setting SLACK_POST_WEBHOOK_URL")
		os.Exit(1)
	}

	var message_buf string
	if len(os.Args) == 2 && exists(os.Args[1]) {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		message_buf = *(*string)(unsafe.Pointer(&b))
	} else {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		message_buf = *(*string)(unsafe.Pointer(&b))
	}

	params, err := json.Marshal(slack{
		fmt.Sprintf("```%s```", message_buf),
		slack_username,
		slack_icon,
		slack_post_channel})
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.PostForm(
		webhook_url,
		url.Values{"payload": {string(params)}},
	)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	log.Println(string(body))
}
