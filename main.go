package btbgo

import (
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
	func speedType(AuthToken string) string{
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
		return string(body)


	}
func madlibs(AuthToken string) string{
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
	return string(body)


}
func lyrics(AuthToken string, song string) string{
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
	return string(body)


}
func meme(AuthToken string) string{
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
	return string(body)


}
func word(AuthToken string) string{
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
	return string(body)


}
func reddit(AuthToken string, query string, limit int) string{
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
	return string(body)

}
