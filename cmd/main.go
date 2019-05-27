package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/r4wm/elastic_kjv"
)

//Create the elasticsearch db
func main() {
	var dbPath = flag.String("dbPath", "", "Path to kjv db file: ../data/kjv.db")
	var filePath = flag.String("out", "/tmp/esBulk.json", "file output path")
	flag.Parse()

	if *dbPath == "" {
		fmt.Printf("Must provide path to dbfile")
		os.Exit(1)
	}

	stuff, err := elastic_kjv.PullFromSQL(*dbPath)
	if err != nil {
		panic(err)
	}

	// fmt.Println(stuff)

	esBulkPayload, err := elastic_kjv.CreateESBulkPost(&stuff)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create(*filePath)
	defer f.Close()
	written, err := esBulkPayload.WriteTo(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes to %s: ", written, *filePath)

}
