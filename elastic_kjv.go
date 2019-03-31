package elastic_kjv

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// ElasticVerse contents for bible verse.
type ElasticVerse struct {
	LinearOrderedVerse   int    `json:"linearOrderedVerse"`
	LinearOrderedChapter int    `json:"linearOrderedChapter"`
	Testament            string `json:"testament"`
	Chapter              int    `json:"chapter"`
	Book                 string `json:"book"`
	Verse                int    `json:"verse"`
	Text                 string `json:"text"`
}

// PullFromSQL pulls datas from sql and returns array of ElasticVerse objects
func PullFromSQL() ([]ElasticVerse, error) {
	dbFile := "/home/rmintz/kjv.db"
	bulk := []ElasticVerse{}

	// Connect
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Println("Error PullFromSQL cant open db")

		return nil, err
	}

	// Query
	rows, err := db.Query("select * from kjv;")
	if err != nil {
		fmt.Println("failed to query db.")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		singleVerse := ElasticVerse{}
		err = rows.Scan(
			&singleVerse.Book,
			&singleVerse.Chapter,
			&singleVerse.Verse,
			&singleVerse.Text,
			&singleVerse.LinearOrderedVerse,
			&singleVerse.LinearOrderedChapter,
			&singleVerse.Testament,
		)

		if err != nil {
			fmt.Println("Failed to parse query.")
			return nil, err
		}

		bulk = append(bulk, singleVerse)
	}

	return bulk, nil

}

// CreateESBulkPost returns a newline delimited json string to POST into ES.
func CreateESBulkPost(ebulk *[]ElasticVerse) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	type ESIndex struct {
		Index struct {
			Index string `json:"_index"`
			Type  string `json:"_type"`
			ID    int    `json:"_id"`
		} `json:"index"`
	}

	// Create the string buffer
	for _, v := range *ebulk {
		singleItem := ESIndex{}
		singleItem.Index.Index = "bible"
		singleItem.Index.Type = "kjv"
		singleItem.Index.ID = v.LinearOrderedVerse

		// fmt.Printf("%#v\n", singleItem)
		result, _ := json.Marshal(singleItem)
		result = append(result, '\n')
		buffer.Write(result)

		// Handle body
		body, _ := json.Marshal(v)
		body = append(body, '\n')
		buffer.Write(body)
	}
	// ES Bulk require newline termination
	buffer.Write([]byte("\n"))

	return buffer, nil
}
