package main

import (
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func getEmoji() <-chan string {
	emojiCh := make(chan string)

	go func() {
		defer close(emojiCh)
		for {
			emojiSets := [][]int{
				// Emoticons icons.
				{128513, 128591},
				// Dingbats.
				{9986, 10160},
				// Transport and map symbols.
				{128640, 128704},
			}
			// access random emoji set
			setIndex := rand.Intn(3)
			set := emojiSets[setIndex]
			// get random emoji in set
			emojiCode := rand.Intn(set[1]-set[0]) + set[0]
			emoji := html.UnescapeString("&#" + strconv.Itoa(emojiCode) + ";")
			emojiCh <- emoji
		}

	}()

	return emojiCh
}

func emojiHandler(emojiCh <-chan string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		emoji := <-emojiCh
		log.Printf("Sending emoji %s\n", emoji)
		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(200)
		fmt.Fprintf(w, "Your emoji is %s\n", emoji)
	}
}

func main() {
	emojiCh := getEmoji()
	log.Fatal(http.ListenAndServe(":8080", emojiHandler(emojiCh)))
}

func init() {
	rand.Seed(time.Now().Unix())
}
