package cosmosapp

type ValidatorCommissionResp struct {
	Commissions ValidatorCommissions `json:"commission"`
}

type ValidatorCommissions struct {
	Commission []ValidatorCommission `json:"commission"`
}

type ValidatorCommission struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
