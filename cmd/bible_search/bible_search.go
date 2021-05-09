package main

import (
	"flag"
	"fmt"

	"github.com/r4wm/elastic_kjv/query"
)

func main() {
	var resultSize = flag.Int64("size", 10, "set results count limit")
	var jsonOut = flag.Bool("json", false, "output json")
	flag.Parse()

	result := query.MakeAndQuery(flag.Args()[:], *resultSize, *jsonOut)
	fmt.Println(string(result))
}
