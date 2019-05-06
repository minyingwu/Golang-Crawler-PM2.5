package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"Go/pm2.5/models"
)

var p models.Pollution

type serverResponse struct {
	res     *http.Response
	err     error
	retries int
}

func newServerResponse(retry int) *serverResponse {
	serverRes := new(serverResponse)
	serverRes.retries = retry
	return serverRes
}

func helpRead(res *http.Response) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(body, &p.DataSlice)
	log.Println("Loading finish")
	insertPollutionDB()
}

func insertPollutionDB() {
	p.RemoveAllPollution()
	p.InsertPollution()
	log.Println("Save DB finish")
	if !isServerOn {
		ch <- true
		close(ch)
	}
}

// getAirPollutionData request new pollution data from open server
func getAirPollutionData() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for ; true; <-ticker.C {
		serRes := newServerResponse(3)
		// Retry getting data
		for serRes.retries > 0 {
			serRes.res, serRes.err = http.Get("https://opendata.epa.gov.tw/ws/Data/ATM00625/?$format=json")
			if serRes.err != nil {
				log.Println(serRes.err)
				serRes.retries--
				time.Sleep(3 * time.Second)
			} else {
				break
			}
		}

		if serRes.res != nil {
			helpRead(serRes.res)
		} else {
			fmt.Println("Server response is empty !!")
		}
	}
}
