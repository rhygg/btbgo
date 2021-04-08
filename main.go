package btbgo

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
	func SpeedType(AuthToken string) []byte{
		var url string = "https://api.bytestobits.dev/text"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
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
func Madlibs(AuthToken string) []byte{
	var url string = "https://api.bytestobits.dev/madlibs"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
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
func Lyrics(AuthToken string, song string) []byte{
	var url string = "https://api.bytestobits.dev/lyrics+song="+song
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
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
func Meme(AuthToken string) []byte{
	var url string = "https://api.bytestobits.dev/meme"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
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
func Word(AuthToken string) []byte{
	var url string = "https://api.bytestobits.dev/reddit"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
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
func Reddit(AuthToken string, query string, limit int) []byte{
	var url string = "https://api.bytestobits.dev/reddit"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
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
