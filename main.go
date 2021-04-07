package btbgo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

/*func main() {
	var token string = "oO17.WHFrSanuzvxVwFXljBW4"
    var song string= "What's+next"
	log.Println(speedType(token))
	log.Println(madlibs(token))
}

 */
	func speedType(AuthToken string) []byte{
		var url string = "https://api.bytestobits.dev/text"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("authorization", AuthToken)
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
func madlibs(AuthToken string) []byte{
	var url string = "https://api.bytestobits.dev/madlibs"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("authorization", AuthToken)
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
func lyrics(AuthToken string, song string) []byte{
	var url string = "https://api.bytestobits.dev/lyrics+song="+song
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("authorization", AuthToken)
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
func meme(AuthToken string) []byte{
	var url string = "https://api.bytestobits.dev/meme"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("authorization", AuthToken)
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
func word(AuthToken string) []byte{
	var url string = "https://api.bytestobits.dev/reddit"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("authorization", AuthToken)
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
func reddit(AuthToken string, query string, limit int) []byte{
	var url string = "https://api.bytestobits.dev/reddit"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("authorization", AuthToken)
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
