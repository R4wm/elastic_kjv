package main

import (
	"fmt"
	"os"

	"github.com/r4wm/elastic_kjv"
)

//Create the elasticsearch db
func main() {

	filePath := "/tmp/esBulk.json"

	stuff, err := elastic_kjv.PullFromSQL()
	if err != nil {
		panic(err)
	}

	// fmt.Println(stuff)

	esBulkPayload, err := elastic_kjv.CreateESBulkPost(&stuff)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create(filePath)
	defer f.Close()
	written, err := esBulkPayload.WriteTo(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes to %s: ", written, filePath)

}
