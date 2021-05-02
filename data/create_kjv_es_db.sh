#!/bin/bash

index="kjv"

echo "setting search as you type"
curl -X PUT "localhost:9200/$index?pretty" -H 'Content-Type: application/json' -d'
{
  "mappings": {
    "properties": {
      "verse": {
        "type": "search_as_you_type"
      }
    }
  }
}
'

echo "Setting replica count to 1"
curl -X PUT "localhost:9200/$index/_settings?pretty" -H 'Content-Type: application/json' -d'
{
  "index" : {
    "number_of_replicas" : 1
  }
}
'


curl -XPOST -H"Content-Type: application/json" http://localhost:9200/kjv/_bulk --data-binary @es7Bulk.json

