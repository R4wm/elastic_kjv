# bible_search

TL;DR: A quick way to search keywords in the bible

A command line tool to query ES bible index
- json out or plain text
- set the size of results you want
  - so for "top 3" `-size 3`


### Usage
```bash

$ go run bible_search.go -h
Usage of /tmp/go-build793390229/b001/exe/bible_search:
  -json
    	output json
  -size int
    	set results count limit (default 10)

```

### json output
```json

00:35 $ go run bible_search.go -size 1 -json mystery   | jq . 
{
  "hits": [
    {
      "_id": "29741",
      "_index": "bible",
      "_score": 9.619158,
      "_source": {
        "book": "1TIMOTHY",
        "chapter": 3,
        "linearOrderedChapter": 54,
        "linearOrderedVerse": 29741,
        "testament": "NEW",
        "text": "Holding the mystery of the faith in a pure conscience.",
        "verse": 9
      },
      "_type": "_doc"
    }
  ],
  "max_score": 9.619158,
  "total": {
    "relation": "eq",
    "value": 22
  }
}
✔ ~/go/src/github.com/r4wm/elastic_kjv/cmd/bible_search [json_output L|✚ 1]
```


### plain text output
```bash

✔ ~/go/src/github.com/r4wm/elastic_kjv [master|✔] 
09:37 $ bible_search -size 5 gospel kingdom grace
GALATIANS 1:6
 I marvel that ye are so soon removed from him that called you into the grace of Christ unto another gospel:

MARK 1:14
 Now after that John was put in prison, Jesus came into Galilee, preaching the gospel of the kingdom of God,

MARK 1:15
 And saying, The time is fulfilled, and the kingdom of God is at hand: repent ye, and believe the gospel.

MATTHEW 24:14
 And this gospel of the kingdom shall be preached in all the world for a witness unto all nations; and then shall the end come.

HEBREWS 12:28
 Wherefore we receiving a kingdom which cannot be moved, let us have grace, whereby we may serve God acceptably with reverence and godly fear:

```

```bash
✔ ~/go/src/github.com/r4wm/elastic_kjv [master|✔] 
09:38 $ bible_search -size 3 gospel thy word
1PETER 1:25
 But the word of the Lord endureth for ever. And this is the word which by the gospel is preached unto you.

COLOSSIANS 1:5
 For the hope which is laid up for you in heaven, whereof ye heard before in the word of the truth of the gospel;

ACTS 8:25
 And they, when they had testified and preached the word of the Lord, returned to Jerusalem, and preached the gospel in many villages of the Samaritans.

✔ ~/go/src/github.com/r4wm/elastic_kjv [master|✔] 
09:38 $ 

```