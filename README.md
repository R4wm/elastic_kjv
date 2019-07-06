# elastic_kjv
Creates elasticsearch db from sqlite db

Get it. Note that r4wm is not Capitalized.. 
```bash
go get -v -u github.com/r4wm/elastic_kjv
```


Check it!
```bash
ᚱ@data $ md5sum -c md5sums 
create_kjv_es_db.sh: OK
es6Bulk.json: OK
es7Bulk.json: OK
kjv.db: OK
ᚱ@data $ 
```


Create the payload
```bash
go run main.go -dbPath ../data/kjv.db -out ../data/es7Bulk.json
```


Post in the payload.

```bash
ᚱ@cmd $ curl -XPOST -H"Content-type: application/json" http://localhost:9200/_bulk --data-binary @/home/rmintz/go/src/github.com/r4wm/elastic_kjv/data/es7Bulk.json #json output removed.
ᚱ@cmd $ curl -s http://localhost:9200/_cat/indices
yellow open bible CTl0JD1BTIulTVfYcJksDA 1 1 31102 0 12.6mb 12.6mb
ᚱ@cmd $ 
ᚱ@cmd $ 
ᚱ@cmd $ curl -s http://localhost:9200/bible/_search?size=1 | jq . 
{
  "took": 69,
  "timed_out": false,
  "_shards": {
    "total": 1,
    "successful": 1,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": {
      "value": 10000,
      "relation": "gte"
    },
    "max_score": 1,
    "hits": [
      {
        "_index": "bible",
        "_type": "_doc",
        "_id": "1",
        "_score": 1,
        "_source": {
          "linearOrderedVerse": 1,
          "linearOrderedChapter": 1,
          "testament": "OLD",
          "chapter": 1,
          "book": "GENESIS",
          "verse": 1,
          "text": "In the beginning God created the heaven and the earth."
        }
      }
    ]
  }
}
```