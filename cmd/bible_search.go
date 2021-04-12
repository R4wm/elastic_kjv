package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {

	args := os.Args
	stringArg := strings.Join(args[1:], " ")

	es, _ := elasticsearch.NewDefaultClient()
	// log.Println(elasticsearch.Version)
	// log.Println(es.Info())

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"text": stringArg,
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

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		book := hit.(map[string]interface{})["_source"].(map[string]interface{})["book"]
		text := hit.(map[string]interface{})["_source"].(map[string]interface{})["text"]
		chapter := hit.(map[string]interface{})["_source"].(map[string]interface{})["chapter"]
		verse := hit.(map[string]interface{})["_source"].(map[string]interface{})["verse"]
		fmt.Printf("%s %d:%d\n %s\n\n", book, int(chapter.(float64)), int(verse.(float64)), text)
		// log.Printf("* ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
}
