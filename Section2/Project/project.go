package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    18.00,
	"BOOK":   25.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrices[itemCode]

	if found {
		return basePrice, true
	}

	if strings.HasSuffix(itemCode, "_SALE") {
		originalItemCode := strings.TrimSuffix(itemCode, "_SALE")

		basePrice, found = productPrices[originalItemCode]

		if found {
			salePrice := basePrice * 0.90
			return salePrice, true
		}
	}

	fmt.Printf("%s not found\n", itemCode)
	return 0.0, false
}

func main() {
	price, found := calculateItemPrice("TSHIRT_SALE")

	if found {
		fmt.Println("Price:", price)
	}
}
