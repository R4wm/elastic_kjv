package query

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

// returns query_string: "text: just AND text: faith AND text: alone"
func GetQueryString(keywords []string) string {
	var b strings.Builder

	if len(keywords) == 0 {
		return ""
	}

	lastKeywordNum := (len(keywords) - 1)
	for i, v := range keywords {
		b.WriteString(fmt.Sprintf("text: %s", v))
		if i != lastKeywordNum {
			b.WriteString(" AND ")
		}
	}

	return b.String()
}

func PlainOut(hit map[string]interface{}) []byte {

	var buf bytes.Buffer
	for _, hit := range hit["hits"].(map[string]interface{})["hits"].([]interface{}) {
		book := hit.(map[string]interface{})["_source"].(map[string]interface{})["book"]
		text := hit.(map[string]interface{})["_source"].(map[string]interface{})["text"]
		chapter := hit.(map[string]interface{})["_source"].(map[string]interface{})["chapter"]
		verse := hit.(map[string]interface{})["_source"].(map[string]interface{})["verse"]

		buf.WriteString(fmt.Sprintf("%s %d:%d\n %s\n\n", book, int(chapter.(float64)), int(verse.(float64)), text))
	}
	return buf.Bytes()
}

func MakeAndQuery(aWords []string, aSize int64, aJson bool) []byte {
	var buf bytes.Buffer

	es, _ := elasticsearch.NewDefaultClient()
	parsedQueryString := GetQueryString(aWords)
	log.Print("parsedQuerystring: ", parsedQueryString)

	query := map[string]interface{}{
		"size": aSize,
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"default_field": "text",
				"query":         GetQueryString(aWords),
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
		log.Print("Failed to get json decode response from ES")
	}

	jsonResult, err := json.Marshal(r["hits"])
	if err != nil {
		log.Fatalf("Failed to parse es query result: %v\n", err)
	}
	if aJson {
		return jsonResult
	}

	return PlainOut(r)

}
