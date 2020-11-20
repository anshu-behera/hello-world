package main

import (
	"os"
	"github.com/gorilla/mux"
	"time"
	"io/ioutil"
	"log"
	"net/http"

)

func main() {
	log.Println("help")
	client := http.Client{}

	r := mux.NewRouter()
	myHandler := func (w http.ResponseWriter, r *http.Request)  {
		resp, _ := client.Get("https://jsonplaceholder.typicode.com/todos/4")
		body, _ := ioutil.ReadAll(resp.Body)
		w.Write(body)
	}

	myHandler2 := func (w http.ResponseWriter, r *http.Request)  {
		os.Exit(0)
	}
	r.HandleFunc("/",myHandler)
	r.HandleFunc("/exit",myHandler2 )
	s := &http.Server{
		Addr:           ":8085",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
