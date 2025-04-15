package models

type PaymentMethod struct {
	Account string `json:"account"`
	Status  string `json:"status"`
	Amount string `json:"amount"`
	Icon    string `json:"icon"`
}
