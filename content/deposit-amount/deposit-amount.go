package deposit_amount

type Request struct {
	Token         string `json:"token"`
	DepositAmount string `json:"deposit_amount"`
}

type Response struct {
	TempToken   string `json:"temp_token"`
	TotalAmount string `json:"total_amount"`
}
