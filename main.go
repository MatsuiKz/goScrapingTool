package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 1. ウェブページの取得
	res, err := http.Get("https://news.yahoo.co.jp/rss/topics/top-picks.xml")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// 2. HTMLの解析
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	// 3. 必要な情報の抽出
	title := doc.Find("title").Text()
	pubDate := doc.Find("pubDate").Text()
	fmt.Println(title)
	fmt.Println(pubDate)
}
