package ecsite

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Amazon struct {
	url string
}

func NewAmazon(url string) *Amazon {
	return &Amazon{
		url: url,
	}
}

func (a *Amazon) JudgeAmazon() (string, error) {
	// fmt.Printf((a.url))
	doc, err := goquery.NewDocument(a.url)
	if err != nil {
		fmt.Println(err.Error())
		return "error", err
	}
	e := doc.Find("#addToCart_feature_div")

	if e.Length() > 0 {
		return a.url, nil
	}
	return "在庫なし", nil


}
