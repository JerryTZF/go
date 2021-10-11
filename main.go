/**
 * Created by GoLand
 * Time: 2021/9/7 1:58 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: main.go
 * Desc:
 */
package main

import (
	"github.com/PuerkitoBio/goquery"
	"golang-study/pkg"
	"io"
	"log"
	"net/http"
)

func main() {
	//library.StrIndex()
	pkg.Main()
	//library.Json2Struct()
	//demo()
}

type Record struct {
	Currency     string  // 币种
	XianHuiBuy   float64 // 现汇买入
	XianChaoBuy  float64 // 现钞买入
	XianHuiSell  float64 // 现汇卖出
	XianChaoSell float64 // 现钞卖出
}

type Bank struct {
	Name         string
	Date         string
	ExchangeRate []Record
}

var banks = []string{"cgb", "cib", "ccb", "cmb"}

func demo() {
	res, err := http.Get("https://www.kylc.com/bank/rmbfx/b-cib.html")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//rs := make([]Record, 10)

	doc.Find("tbody tr").Each(func(i int, selection *goquery.Selection) {

	})

}
