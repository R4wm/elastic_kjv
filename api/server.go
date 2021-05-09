package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/r4wm/elastic_kjv/query"
)

func esWordsSearch(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v\n", err)
	}

	queryStr := r.FormValue("q")
	if len(queryStr) <= 0 {
		w.Write([]byte("i think your missing search argument.."))
		return
	}

	requestedJson := r.FormValue("json")
	var resultSize int64

	resultSize, err := strconv.ParseInt(r.FormValue("size")[0:], 10, 64)
	if err != nil {
		resultSize = 10
	}

	jsonResult := query.MakeAndQuery(strings.Split(queryStr, " "), resultSize, requestedJson == "true")

	w.Write([]byte(jsonResult))
}

func main() {
	port := "127.0.0.1:8081"

	http.HandleFunc("/words_search", esWordsSearch)

	fmt.Printf("Starting web server")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
