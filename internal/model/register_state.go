package model

type RegisterState struct {
	ID              int `json:"id"`
	FiscalYear      int `json:"fiscalYear"`
	FiscalReceipt   int `json:"fiscalReceipt"`
	BusinessReceipt int `json:"businessReceipt"`
}
