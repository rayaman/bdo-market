package data

import (
	"fmt"

	"github.com/rayaman/bdo-market/app/types"
)

// Filter filters the items using the provided filter function.
//
// Parameters:
// - items []types.MarketItem
// - filter func(types.MarketItem) bool
// Return type: []types.MarketItem
func Filter(items []types.MarketItem, filter func(types.MarketItem) bool) []types.MarketItem {
	var filteredItems []types.MarketItem
	for _, item := range items {
		if filter(item) {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}

func SprintMarketItem(item types.MarketItem) string {
	return fmt.Sprintf("Name: %s\n    Id: %d\n    Current Stock: %d\n    Total Trades: %d\n    Base Price: %d\n    Main Category: %d\n    Sub Category: %d\n    Market Category: %d", item.Name, item.Id, item.CurrentStock, item.TotalTrades, item.BasePrice, item.MainCategory, item.SubCategory, item.MarketCategory)
}
