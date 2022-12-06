package luckyBag

import (
	"math/rand"
	"time"

	"github.com/szpp-dev-team/szpp-slack-bot/model"
	"github.com/szpp-dev-team/szpp-slack-bot/scrape"
)

func MakeLuckyBag(price int) ([]model.Product, error) {

	scrapeResult, err := scrape.Scrape()
	if err != nil {
		return nil, err
	}

	products := scrapeResult.Products
	len := len(products)
	table := make([]int, len)
	for i := range table {
		table[i] = 0
	}
	cnt := 0
	sum := 0

	for true {
		n := makeRand(len)
		if table[n] != 0 { // すでにその商品を確認済みならやり直し
			continue
		}

		if sum+products[n].Price < price { // 合計金額を越えなければ追加する
			table[n] = 1
			sum += products[n].Price
		} else {
			table[n] = -1
		}

		cnt++

		if cnt >= len || float64(price)*0.9 < float64(sum) { // 全ての商品を見るか、合計金額が最大値の9割を超えたら終了
			break
		}
	}

	var resp []model.Product
	for i, row := range table {
		if row == 1 {
			resp = append(resp, products[i])
		}
	}

	return resp, nil
}

func makeRand(max int) int {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	return r.Intn(max)
}
