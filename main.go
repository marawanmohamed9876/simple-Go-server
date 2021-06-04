package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Point struct {
	x int
	y int
}

func requestHandler(response http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(response, "<h1> Hello, world!</h1>")
	fmt.Println("Started!!!!")
	body, _ := ioutil.ReadAll(request.Body)
	var point Point
	err := json.Unmarshal(body, &point)
	if err != nil {
		panic(err)
	}
	log.Printf("the recieved point is X=%d Y=%d", point.x, point.y)
	fmt.Fprintf(response, "received!")
}

func main() {
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":8050", nil)

	if err != nil {
		panic(err)
	}

}
