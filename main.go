package main

import (
	"app/types"
	"fmt"
	"json"
	"net/http"
)

func main() {

}

func GetMarketItems() []types.MarketItem {
	var marketdata []types.MarketItem
	url := "https://api.arsha.io/v2/na/market"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	json.UnMarshal(res.Body, &marketdata)
	return marketdata
}
