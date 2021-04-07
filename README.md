# Bytes To Bits Api wrapper for golang

## How it works

**Go** over to [the bytestobits api](api.bytestobits.dev/) page and [create](api.bytestobits.dev/account) an account, and get an api authentication token.

### What's next?
**Well** use it for your golang projects!

**A pretty example of how you can use the wrapper**

```go
package main
import(
"log"
"github.com/rhydderchc/btbgo"
)
func main(){
var authToken string = "YOUR AUTH_TOKEN"
log.Println(btbgo.text(authToken))
log.Println(btbgo.madlibs(authToken))
}
```
## Endpoints
o /meme -> `btb.meme`

o /madlibs -> `btb.madlibs`

o /text -> `btb.text`

o /word -> `btb.word`


**Un implemented yet**
o /lyrics -> `btb.lyrics`

o /subreddit -> `btb.subreddit`

## TODOS

[X] Complete the lyrics and reddit endpoints
