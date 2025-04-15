package model

type Business struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	City       string `json:"city"`
	Address    string `json:"address"`
	BusinessID string `json:"businessId"`
	TaxID      string `json:"taxId"`
}
