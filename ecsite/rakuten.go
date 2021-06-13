package ecsite

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const ps5Digital = "https://books.rakuten.co.jp/rb/16462860/"
const ps5Normal = "https://books.rakuten.co.jp/rb/16462859/"

type Rakuten struct {
	url string
}

func NewRakuten(url string) *Rakuten {
	return &Rakuten{
		url: url,
	}
}

func (r *Rakuten) JudgeRakuten() (string, error) {
	doc, err := goquery.NewDocument(r.url)
	if err != nil {
		fmt.Println(err.Error())
		return "error", err
	}

	e := doc.Find("body > div.dui-container.main > div.dui-container.content > div.dui-container.searchresults > div > div:nth-child(1) > div.content.status > span")
	if e.Text() != "売り切れ" {
		fmt.Println("デジタル版")
		return ps5Digital, nil
	}
	e = doc.Find("body > div.dui-container.main > div.dui-container.content > div.dui-container.searchresults > div > div:nth-child(2) > div.content.status > span")
	if e.Text() != "売り切れ" {
		fmt.Println("通常版")
		return ps5Normal, nil
	}

	return "在庫なし", nil

}
