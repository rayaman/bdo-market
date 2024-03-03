package types

// {"name":"Lightstone of Flora: Haggler","id":764021,"currentStock":5,"totalTrades":66994,"basePrice":5300000,"mainCategory":85,"subCategory":4}
type MarketItem struct {
	Name           string `json:"name"`
	Id             int    `json:"id"`
	CurrentStock   int    `json:"currentStock"`
	TotalTrades    int    `json:"totalTrades"`
	BasePrice      int    `json:"basePrice"`
	MainCategory   int    `json:"mainCategory"`
	SubCategory    int    `json:"subCategory"`
	MarketCategory int    `json:"marketCategory"`
}
