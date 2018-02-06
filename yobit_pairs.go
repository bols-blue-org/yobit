package yobit_bot

import (
	"encoding/json"

	yobit "github.com/bols-blue-org/go-yobit"
)

type CoinPairList struct {
	Pairs      map[string]coin_basic `json:"pairs"`
	ServerTime int                   `json:"server_time"`
}

type coin_basic struct {
	DecimalPlaces int     `json:"decimal_places"`
	Fee           float64 `json:"fee"`
	FeeBuyer      float64 `json:"fee_buyer"`
	FeeSeller     float64 `json:"fee_seller"`
	Hidden        int     `json:"hidden"`
	MaxPrice      int     `json:"max_price"`
	MinAmount     float64 `json:"min_amount"`
	MinPrice      int     `json:"min_price"`
	MinTotal      float64 `json:"min_total"`
}

func GetCoinPairList(client yobit.Yobit) (CoinPairList, error) {
	url := "https://yobit.net/api/3/info"
	allCoins := CoinPairList{}
	r, err := client.Do("info", "", "")
	if err != nil {
		return allCoins, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(allCoins)
	return allCoins, err
}
