package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "ooops", http.StatusBadRequest)
			return
		}
		log.Println("name %s ", d)
	})
	http.ListenAndServe(":9090", nil)
}
