package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var url string
var idArray []string

type USER struct {
	Id     int32
	UserId int32
	Title  string
	body   string
}

func main() {

	idArray = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	url = "https://jsonplaceholder.typicode.com/posts/"
	taskCompleted := 0

	for _, id := range idArray {

		go func(id string) {

			fmt.Println("Welcome, id is " + id)

			resp, err := http.Get(url + id)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			taskCompleted++
			fmt.Println(string(body))

		}(id)
	}
	for taskCompleted < len(idArray) {
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("All goroutines completed....")
}
