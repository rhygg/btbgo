package btb

import (
	"io/ioutil"
	"log"
	"net/http"
)

var authTok string

func RegisterToken(auth) string{
	authTok = auth
	return authTok
}
	func Text() []byte{
		var url string = "https://api.bytestobits.dev/text"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		if len(authTok) == 0{
			log.Fatalln("The authentication token was not provided")
			return nil
		}
		req.Header.Set("authorization", authTok)
		res, err := client.Do(req)
		if err!=nil {
			log.Fatalln(err)

		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		textByte :=[]byte(body)

		return textByte


	}
func Madlibs() []byte{
	var url string = "https://api.bytestobits.dev/madlibs"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if len(authTok) == 0{
			log.Fatalln("The authentication token was not provided")
			return nil
		}
	req.Header.Set("authorization", authTok)
	res, err := client.Do(req)
	if err!=nil {
		log.Fatalln(err)

	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	textByte :=[]byte(body)

		return textByte


}
func Lyrics(song string) []byte{
	var url string = "https://api.bytestobits.dev/lyrics+song="+song
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if len(authTok) == 0{
			log.Fatalln("The authentication token was not provided")
			return nil
		}
	req.Header.Set("authorization", authTok)
	res, err := client.Do(req)
	if err!=nil {
		log.Fatalln(err)

	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	textByte :=[]byte(body)

		return textByte


}
func LyricsWithArtist(song string, artist string) []byte{
		var url string = "https://api.bytestobits.dev/lyrics+song="+song+"&artist="+artist
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if len(authTok) == 0{
			log.Fatalln("The authentication token was not provided")
			return nil
		}
	req.Header.Set("authorization", authTok)
	res, err := client.Do(req)
	if err!=nil {
		log.Fatalln(err)

	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	textByte :=[]byte(body)

		return textByte		
	
}
func Meme() []byte{
	var url string = "https://api.bytestobits.dev/meme"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if len(authTok) == 0{
			log.Fatalln("The authentication token was not provided")
			return nil
		}
	req.Header.Set("authorization", authTok)
	res, err := client.Do(req)
	if err!=nil {
		log.Fatalln(err)

	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	textByte :=[]byte(body)

		return textByte


}
func Word() []byte{
	var url string = "https://api.bytestobits.dev/reddit"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if len(authTok) == 0{
			log.Fatalln("The authentication token was not provided")
			return nil
		}
	req.Header.Set("authorization", authTok)
	res, err := client.Do(req)
	if err!=nil {
		log.Fatalln(err)

	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	textByte :=[]byte(body)

		return textByte


}
func Reddit(query string) []byte{
	var url string = "https://api.bytestobits.dev/reddit?subreddit="+query
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if len(authTok) == 0{
			log.Fatalln("The authentication token was not provided")
			return nil
		}
	req.Header.Set("authorization", authTok)
	res, err := client.Do(req)
	if err!=nil {
		log.Fatalln(err)

	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	textByte :=[]byte(body)
	return textByte

}
