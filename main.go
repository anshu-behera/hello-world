package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("help")
	client := http.Client{}
	resp, _ := client.Get("https://jsonplaceholder.typicode.com/todos/1")
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
