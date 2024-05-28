package cosmosapp

type SupplyResp struct {
	AmountResp Amount `json:"amount"`
}

type Amount struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}
