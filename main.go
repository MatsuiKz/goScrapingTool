package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	a1 = 123
	fmt.Println(a1)
	// 1. ウェブページの取得
	res, err := http.Get("http://example.com")
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
	fmt.Println(title)
}
