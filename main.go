package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {
	url := "https://codezine.jp/"

	c := colly.NewCollector()

	t := time.Now()
	f, err := os.Create("./dailyRanking/"+ t.Format("20060102") +".txt")
	
	//中身を空にする
	f.Truncate(0)

	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	
	i := 1
	c.OnHTML(".c-primarysection_body .c-ranking_item", func(e *colly.HTMLElement) {
		if i > 10 {
			return
		}
		fmt.Printf("do %d \n" , i)
		
		ranking := e.DOM.Find("dl > dt").Text()
		title := e.DOM.Find("dl > dd > a").Text()
		linq := e.DOM.Find("dl > dd > a").AttrOr("href", "")
		str := fmt.Sprintf("%s位 \n タイトル: %s\n URL: %s \n\n", ranking, title, url + linq)
		str = strings.Replace(str, "　", " ", -1)
		f.Write([]byte(str))
		i++
	})


	c.Visit(url)
}