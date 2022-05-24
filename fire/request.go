package fire

import (
	"io/ioutil"
	"log"
	"net/http"
)

var (
	INSTRUMENT_LIST_URI = "https://margincalculator.angelbroking.com/OpenAPI_File/files/OpenAPIScripMaster.json"
)

//executes all request placing the auth headers
func fireRequest(req *http.Request) ([]byte, int) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, res.StatusCode
}

func getInstrumentListRequest() *http.Request {
	var URI = INSTRUMENT_LIST_URI
	req, _ := http.NewRequest("GET", URI, nil)
	return req
}

func GetInstruments() ([]byte, int) {
	req := getInstrumentListRequest()
	body, status := fireRequest(req)
	return body, status
}
