package scrape

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestScrape(t *testing.T) {
	result, err := Scrape()
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range result.Products {
		fmt.Println(strconv.Itoa(i) + ": ")
		fmt.Println(row.Name)
		fmt.Println(row.Price)
		fmt.Println(row.ThumbnailImageURL)
		if row.IsPrime {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not Prime")
		}
	}
}
