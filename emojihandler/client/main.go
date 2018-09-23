package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		resp, err := http.Get("http://localhost:8080")
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(string(data))
		time.Sleep(1 * time.Second)
	}
}
