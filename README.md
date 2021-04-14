

# elastic_kjv
Creates elasticsearch db from sqlite db
If you dont have the kjv sqlite db, you will need that. Dont worry, its easy: see [bible_api](https://github.com/R4wm/bible_api/blob/master/cmd/deploy.go#L19)

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
### bible_search cli
Command line tool to perform query_string searches. 
Find it [here](https://github.com/R4wm/elastic_kjv/blob/master/cmd/bible_search/bible_search.go)
```bash
✔ ~/go/src/github.com/r4wm/elastic_kjv [master|✔] 
00:15 $ go build -o ~/bin/ cmd/bible_search/bible_search.go 

✔ ~/go/src/github.com/r4wm/elastic_kjv [master|✔] 
00:15 $ time bible_search -size 4 gospel kingdom 
MARK 1:14
 Now after that John was put in prison, Jesus came into Galilee, preaching the gospel of the kingdom of God,

MARK 1:15
 And saying, The time is fulfilled, and the kingdom of God is at hand: repent ye, and believe the gospel.

MATTHEW 24:14
 And this gospel of the kingdom shall be preached in all the world for a witness unto all nations; and then shall the end come.

MATTHEW 9:35
 And Jesus went about all the cities and villages, teaching in their synagogues, and preaching the gospel of the kingdom, and healing every sickness and every disease among the people.


real	0m0.016s
user	0m0.005s
sys	0m0.009s


✔ ~/go/src/github.com/r4wm/elastic_kjv [master|✔] 
00:15 $ time bible_search -size 4 gospel grace
GALATIANS 1:6
 I marvel that ye are so soon removed from him that called you into the grace of Christ unto another gospel:

PHILIPPIANS 1:7
 Even as it is meet for me to think this of you all, because I have you in my heart; inasmuch as both in my bonds, and in the defence and confirmation of the gospel, ye all are partakers of my grace.

ACTS 20:24
 But none of these things move me, neither count I my life dear unto myself, so that I might finish my course with joy, and the ministry, which I have received of the Lord Jesus, to testify the gospel of the grace of God.


real	0m0.016s
user	0m0.001s
sys	0m0.013s
✔ ~/go/src/github.com/r4wm/elastic_kjv [master|✔] 
00:15 $ 


```

