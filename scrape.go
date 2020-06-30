package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {

	scraper()
	// Instantiate default collector
	// c := colly.NewCollector()

	// // On every a element which has href attribute call callback
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")

	// 	// Print link
	// 	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	// })

	// // c.OnHTML("tr td:nth-of-type(2)", func(e *colly.HTMLElement) {
	// // 	// fmt.Println("First column of a table row:", e.Text)
	// // 	price := e.Attr("td:nth-of-type(3)")

	// // 	fmt.Printf("Name and Price found: %q -> %s\n", e.Text, price)
	// // })

	// // Before making a request print "Visiting ..."
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	// // Start scraping on https://finance.yahoo.com/cryptocurrencies
	// c.Visit("https://finance.yahoo.com/cryptocurrencies")
}

func scraper() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })

	c.OnHTML("tr td:nth-of-type(2)", func(e *colly.HTMLElement) {
		fmt.Println("The name of the Cryptocurrency is:", e.Text)
	})

	c.OnHTML("tr td:nth-of-type(3)", func(e *colly.HTMLElement) {
		fmt.Println("The price of the Cryptocurrency is:", e.Text)
	})

	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://finance.yahoo.com/cryptocurrencies")
}
