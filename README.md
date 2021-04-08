# Bytes To Bits Api wrapper for golang

Download

```
go get github.com/rhydderchc/btbgo
```

## How it works

**Go** over to [the bytestobits api](api.bytestobits.dev/) page and [create](api.bytestobits.dev/account) an account, and get an api authentication token.

## Features
One time authentication!

use `RegisterToken` once and then use all the api functions without any problem!


use all the apis for incredible development!

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
btbgo.registerToken(token)
	err := json.Unmarshal(btbgo.Madlibs(), &getTitle)
	if err != nil {
		log.Println(err)
		return
	}
log.Println(getTitle.Title)
}
```
Returns **the title of the madlibs game**
```Personal Ad```
## Endpoints
 /meme -> `btb.Meme`

 /madlibs -> `btb.Madlibs`

 /text -> `btb.Text`

 /word -> `btb.Word`


**Un-implemented yet**

 /lyrics -> `btb.Lyrics`

 /subreddit -> `btb.Subreddit`
 

## TODOS

[X] Complete the lyrics and reddit endpoints

# Understanding the wrapper

A very straight forward concept the btbgo wrapper converts the json objects that are returned by http requests and translates them into byte code. This byte code can now be used in any way (the example above) to cover your development with the api.
You can convert the following to string and here's what it would return!
```go
package main
import(
"fmt"
"github.com/rhydderchc/btbgo"
)
func main(){
btbgo.RegisterToken("YOUR_TOKEN")
fmt.Println(string(btbgo.text()))
}
```
Returns
```
"Do you understand the expense people went through to keep you locked up in that tower? You think people like that are just gonna let you walk away? You are an investment and you will not be safe until you are far away from here."
```
Or in native bytes
```go 
package main
import(
"fmt"
"github.com/rhydderchc/btbgo"
)
func main(){
btbgo.RegisterToken("YOUR_TOKEN")
fmt.Println(btbgo.text())
}
```
Returns
```
[34 87 104 97 116 101 118 101 114 32 112 114 105 110 99 105 112 108 101 32 111 102 32 105 110 116 101 108 108 105 103 101 110 99 101 32 119 101 32 97 116 116 97 105 110 32 117 110 116 111 32 105 110 32 116 104 105 115 32 108 105 102 101 44 32 105 116 32 119 105 108 108 32 114 105 115 101 32 119 105 116 104 32 117 115 32 105 110 32 116 104 101 32 114 101 115 117 114 114 101 99 116 105 111 110 46 32 65 110 100 32 105 102 32 97 32 112 101 114 115 111 110 32 103 97 105 110 115 32 109 111 114 101 32 107 110 111 119 108 101 100 103 101 32 97 110 100 32 105 110 116 101 108 108 105 103 101 110 99 101 32 105 110 32 116 104 105 115 32 108 105 102 101 32 116 104 114 111 117 103 104 32 104 105 115 32 100 105 108 105 103 101 110 99 101 32 97 110 100 32 111 98 101 100 105 101 110 99 101 32 116 104 97 110 32 97 110 111 116 104 101 114 44 32 104 101 32 119 105 108 108 32 104 97 118 101 32 115 111 32 109 117 99 104 32 116 104 101 32 97 100 118 97 110 116 97 103 101 32 105 110 32 116 104 101 32 119 111 114 108 100 32 116 111 32 99 111 109 101 46 34 10]

```
