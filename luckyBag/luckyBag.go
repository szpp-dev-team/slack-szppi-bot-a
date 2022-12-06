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
	sum := 0
	var resp []model.Product

	shuffled := shuffleArray(products)
	for _, row := range shuffled {
		if sum + row.Price < price {
			resp = append(resp, row)
			sum += row.Price
		}
	}

	return resp, nil
}

func makeRand(max int) int {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	return r.Intn(max)
}

func shuffleArray(arr []model.Product) []model.Product {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr 
}
