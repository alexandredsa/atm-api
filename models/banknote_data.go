package models

// BanknoteData represents a banknote data with its value and quantity
type BanknoteData struct {
	Quantity int16 `json:"quantity"`
	Value    int16 `json:"value"`
}
