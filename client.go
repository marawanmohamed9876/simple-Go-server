package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Post("http://localhost:8050/",
		"application/json", bytes.NewBuffer([]byte(`{"X":20,"Y":50}`)))

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
