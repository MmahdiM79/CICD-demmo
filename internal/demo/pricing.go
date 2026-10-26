package demo

import "math"

type Item struct {
	Name  string
	Price float64
	Qty   int
}

func CalculateTotal(items []Item, code string) float64 {
	discounts := map[string]float64{"STUDENT20": 0.2, "WELCOME10": 0.1}
	subtotal := 0.0
	for _, item := range items {
		subtotal += item.Price * float64(item.Qty)
	}
	discount := discounts[code]
	return math.Round(subtotal*(1+discount)*100) / 100
}
