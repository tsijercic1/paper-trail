package model

type CreateBusinessRequest struct {
	Name       string `json:"name"`
	City       string `json:"city"`
	Address    string `json:"address"`
	BusinessID string `json:"business_id"`
	TaxID      string `json:"tax_id"`
}
