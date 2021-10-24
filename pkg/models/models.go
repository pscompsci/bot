package models

import "encoding/json"

type Candle struct {
	Pair    string      `json:"pair"`
	EndTime json.Number `json:"endTime"`
	Open    json.Number `json:"open"`
	High    json.Number `json:"high"`
	Low     json.Number `json:"low"`
	Close   json.Number `json:"close"`
	Volume  json.Number `json:"volume"`
}
