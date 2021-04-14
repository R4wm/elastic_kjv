package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/r4wm/elastic_kjv/query"
)

func main() {
	var resultSize = flag.Int("size", 10, "set results count limit")
	var jsonOut = flag.Bool("json", false, "output json")
	flag.Parse()
	// stringArg := strings.Join(flag.Args()[:], " ")
	// query := map[string]interface{}{
	// 	"size": *resultSize,
	// 	"query": map[string]interface{}{
	// 		"match": map[string]interface{}{
	// 			"text": stringArg,
	// 		},
	// 	},
	// }

	es, _ := elasticsearch.NewDefaultClient()
	// log.Println(elasticsearch.Version)
	// log.Println(es.Info())

	var buf bytes.Buffer
	query := map[string]interface{}{
		"size": *resultSize,
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"default_field": "text",
				"query":         query.GetQueryString(flag.Args()[:]),
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("bible"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s\n", err)
	}

	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsint the response body: %s", err)
	}

	// json out
	if *jsonOut {
		jsonResult, _ := json.Marshal(r["hits"])
		fmt.Println(string(jsonResult))
		return
	}

	// plain text out
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		book := hit.(map[string]interface{})["_source"].(map[string]interface{})["book"]
		text := hit.(map[string]interface{})["_source"].(map[string]interface{})["text"]
		chapter := hit.(map[string]interface{})["_source"].(map[string]interface{})["chapter"]
		verse := hit.(map[string]interface{})["_source"].(map[string]interface{})["verse"]
		fmt.Printf("%s %d:%d\n %s\n\n", book, int(chapter.(float64)), int(verse.(float64)), text)
		// log.Printf("* ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
}
