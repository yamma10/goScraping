package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	url := "https://codezine.jp/"

	c := colly.NewCollector()
	
	c.OnHTML(".c-ranking_item", func(e *colly.HTMLElement) {
		ranking := e.DOM.Find("dl > dt").Text()
		title := e.DOM.Find("dl > dd > a").Text()
		linq := e.DOM.Find("dl > dd > a").AttrOr("href", "")
		
		fmt.Printf("ランキング: %s \n タイトル: %s\n URL: %s \n\n", ranking, title, url + linq)
	})

	// c.OnHTML("article", func(e *colly.HTMLElement) {

	// 	i++
	// 	book := e.DOM.Find("a > h3").Text()
	// 	author := e.DOM.Find("div > a").Text()
	// 	if author == "" {
	// 		author = e.ChildText(".BookLink_userName__avtjq")
	// 	}

	// 	fmt.Printf("%d 著者: %s / タイトル: %s\n", i, author, book)
	// })

	c.Visit(url)
}