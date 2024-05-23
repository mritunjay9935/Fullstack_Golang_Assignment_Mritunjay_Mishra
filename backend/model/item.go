package model

type Item struct {
	Key        string      `json:"key"`
	Value      interface{} `json:"value"`
	Expiration int64       `json:"expiration"` // Expiration time in seconds
}
