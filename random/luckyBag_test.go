package luckyBag

import (
	"fmt"
	"log"
	"testing"
)

func TestMakeLuckyBag(t *testing.T) {
	result, err := MakeLuckyBag(10000)
	if err != nil {
		log.Fatal(err)
	}
	
	for i, row := range result {
		fmt.Println(i)
		fmt.Println(row.Name)
		fmt.Println(row.Price)
	}
}
