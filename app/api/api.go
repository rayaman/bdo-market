package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/rayaman/bdo-market/app/types"
)

var items []types.MarketItem
var lastUpdated time.Time = time.Now()

// GetMarketItems retrieves market items from the API and updates them if necessary.
//
// It returns a slice of MarketItem and an error.
func GetMarketItems() ([]types.MarketItem, error) {
	// The api only updates every 30 minutes, so let's only update it every 30 minutes
	if len(items) > 0 && time.Since(lastUpdated) < 30*time.Minute {
		return items, nil
	}
	var marketdata []types.MarketItem

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.arsha.io/v2/na/market", nil)

	if err != nil {
		return marketdata, err
	}

	res, err := client.Do(req)

	if err != nil {
		return marketdata, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return marketdata, err
	}

	json.Unmarshal(body, &marketdata)
	items = marketdata
	return marketdata, nil
}
