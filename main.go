package main

import (
	"fmt"
	"notify/ecsite"
	"notify/line"
	"sync"
)

// normal
const amazonPs5UrlNormal = "https://www.amazon.co.jp/%E3%82%BD%E3%83%8B%E3%83%BC%E3%83%BB%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%A9%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%82%A8%E3%83%B3%E3%82%BF%E3%83%86%E3%82%A4%E3%83%B3%E3%83%A1%E3%83%B3%E3%83%88-PlayStation-5-CFI-1000A01/dp/B08GGGBKRQ"

// digital
const amazonPs5UrlDigital = "https://www.amazon.co.jp/%E3%82%BD%E3%83%8B%E3%83%BC%E3%83%BB%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%A9%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%82%A8%E3%83%B3%E3%82%BF%E3%83%86%E3%82%A4%E3%83%B3%E3%83%A1%E3%83%B3%E3%83%88-PlayStation-5-%E3%83%87%E3%82%B8%E3%82%BF%E3%83%AB%E3%83%BB%E3%82%A8%E3%83%87%E3%82%A3%E3%82%B7%E3%83%A7%E3%83%B3-CFI-1000B01/dp/B08GGF7M7B"

//test
const amazonPs5Test = "https://www.amazon.co.jp/%E3%82%BD%E3%83%8B%E3%83%BC%E3%83%BB%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%A9%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%82%A8%E3%83%B3%E3%82%BF%E3%83%86%E3%82%A4%E3%83%B3%E3%83%A1%E3%83%B3%E3%83%88-%E3%80%90PS5%E3%80%91%E3%83%A9%E3%83%81%E3%82%A7%E3%83%83%E3%83%88-%E3%82%AF%E3%83%A9%E3%83%B3%E3%82%AF-%E3%83%91%E3%83%A9%E3%83%AC%E3%83%AB%E3%83%BB%E3%83%88%E3%83%A9%E3%83%96%E3%83%AB/dp/B08WK1BW23/"

// normaloriginal
const amazonPs5original = "https://www.amazon.co.jp/dp/B08GGGCH3Y/"
// normaloriginaldigital
const amazonPs5OriginalDigital = "https://www.amazon.co.jp/dp/B08GGCGS39/"

// rakuten
const rakutenPs5Url = "https://search.rakuten.co.jp/search/mall/%E6%A5%BD%E5%A4%A9%E3%83%96%E3%83%83%E3%82%AF%E3%82%B9/568376/?f=0"

func main() {
	amazonUrlList := []string{amazonPs5Test, amazonPs5UrlNormal, amazonPs5UrlDigital, amazonPs5original, amazonPs5OriginalDigital}
	c := make(chan string, len(amazonUrlList)+1)
	var wg sync.WaitGroup
	wg.Add(len(amazonUrlList) + 1)

	// Amazon
	for _, amazonUrl := range amazonUrlList {

		go func(url string) {

			a := ecsite.NewAmazon(url)
			message, err := a.JudgeAmazon()
			if err != nil {
				fmt.Errorf("error %v", err)
			}
			c <- message
			wg.Done()

		}(amazonUrl)
	}
	// 楽天
	go func() {
		r := ecsite.NewRakuten(rakutenPs5Url)
		message, err := r.JudgeRakuten()
		if err != nil {
			fmt.Errorf("error %v", err)
		}
		c <- message
		wg.Done()
	}()

	wg.Wait()
	// fmt.Println("終わり")
	// line通知
	var result string
	for i := 0; i < len(amazonUrlList); i++ {
		result = <-c
		if result != "在庫なし" {
			fmt.Println(result)
			l := line.NewLine(result)

			err := l.Notify()
			if err != nil {
				fmt.Errorf("send line error %v", err)
			}

		}
	}


}
