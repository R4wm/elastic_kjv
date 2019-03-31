#!/bin/bash

curl -XPOST -H"Content-Type: application/json" http://localhost:9200/kjv/_bulk --data-binary @esBulk.json

