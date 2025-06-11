package coincap

import "fmt"

type Asset struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[ID] %s \t\t| [Name] %s \t\t| [SYMBOL] %s", d.ID, d.Name, d.Symbol)
}

type assetResponse struct {
	Data []Asset
}
