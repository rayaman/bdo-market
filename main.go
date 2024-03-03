package main

import (
	"fmt"

	"github.com/rayaman/bdo-market/app/api"
	"github.com/rayaman/bdo-market/app/data"
	"github.com/rayaman/bdo-market/app/types"
)

func main() {
	marketdata, err := api.GetMarketItems()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total length: %d\n", len(marketdata))

	filtered_items := data.Filter(marketdata, func(item types.MarketItem) bool {
		return item.CurrentStock > 100000
	})

	fmt.Printf("Filtered length: %d\n", len(filtered_items))
	for _, item := range filtered_items {
		fmt.Println(data.SprintMarketItem(item))
	}
}
