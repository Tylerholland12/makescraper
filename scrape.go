package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/droxey/getpunk/logger"
	"github.com/gocolly/colly"
)

type Stock struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {

	// scraperHandler()

	host := "0.0.0.0:8888"
	http.HandleFunc("/", scraperHandler)
	logger.Log.Info("Server started: http://" + host)

	err := http.ListenAndServe(host, nil)
	if err != nil {
		logger.Log.Error(err)
		return
	}

}

func scraperHandler(w http.ResponseWriter, r *http.Request) { //w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()

	c.OnHTML("body > div.main_content.project > tr td:nth-of-type(2) > tr td:nth-of-type(3) > b", func(e *colly.HTMLElement) {

		stockPrice := new(Stock)
		stockPrice.Name = "name"
		stockPrice.Price = "price"

		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(stockPrice)

		w.Header().Set("Content-Type", "application/json")
		w.Write(bf.Bytes())
	})

	c.Visit("https://finance.yahoo.com/cryptocurrencies")
}
