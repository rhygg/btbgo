# Bytes To Bits Api wrapper for golang

Download

```go get github.com/rhydderchc/btbgo```

## How it works

**Go** over to [the bytestobits api](api.bytestobits.dev/) page and [create](api.bytestobits.dev/account) an account, and get an api authentication token.

### What's next?
**Well** use it for your golang projects!

**A pretty example of how you can use the wrapper**

Say you want to get the title from the **madlibs** endpoint.
```go
package main

import (
	"encoding/json"
	"log"
	"github.com/rhydderchc/btbgo"
)
type MadlibsImplement struct {
	Title string `json:"title"`
}


func main() {
	var token string = "YOUR_API_TOKEN"
getTitle := MadlibsImplement{}
	err := json.Unmarshal(btbgo.Madlibs(token), &getTitle)
	if err != nil {
		log.Println(err)
		return
	}
log.Println(getTitle.Title)
}
```
## Endpoints
o /meme -> `btb.meme`

o /madlibs -> `btb.madlibs`

o /text -> `btb.text`

o /word -> `btb.word`


**Un-implemented yet**

o /lyrics -> `btb.lyrics`

o /subreddit -> `btb.subreddit`

## TODOS

[X] Complete the lyrics and reddit endpoints
