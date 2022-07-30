package discountcalculator

func calculateDiscountRate(clientType string, purchaseValue int) (discountRate float32) {
	if purchaseValue >= 500 || clientType == "gold" {
		return 0.15
	} else if clientType == "silver" || purchaseValue >= 400 {
		discountRate = 0.10
	} else if purchaseValue >= 200 || clientType == "bronze" {
		discountRate = 0.05
	}
	return
}
