package discountcalculator

import "testing"

func Test_calculateDiscountRate(t *testing.T) {
	type args struct {
		clientType    string
		purchaseValue int
	}
	tests := []struct {
		name             string
		args             args
		wantDiscountRate float32
	}{
		{
			"bronze client with purchase value greater equal 199",
			args{
				clientType:    "bronze",
				purchaseValue: 199,
			},
			0.05,
		},
		{
			"silver client with purchase value greater equal 199",
			args{
				clientType:    "silver",
				purchaseValue: 199,
			},
			0.10,
		},
		{
			"gold client with purchase value greater equal 199",
			args{
				clientType:    "gold",
				purchaseValue: 199,
			},
			0.15,
		},
		{
			"bronze client with purchase value greater equal 201",
			args{
				clientType:    "bronze",
				purchaseValue: 201,
			},
			0.05,
		},
		{
			"bronze client with purchase value greater equal 401",
			args{
				clientType:    "bronze",
				purchaseValue: 401,
			},
			0.10,
		},
		{
			"bronze client with purchase value greater equal 400",
			args{
				clientType:    "bronze",
				purchaseValue: 400,
			},
			0.10,
		},
		{
			"bronze client with purchase value greater equal 501",
			args{
				clientType:    "bronze",
				purchaseValue: 501,
			},
			0.15,
		},
		{
			"bronze client with purchase value greater equal 500",
			args{
				clientType:    "bronze",
				purchaseValue: 50,
			},
			0.15,
		},
		{
			"unknown client with purchase value greater equal 501",
			args{
				clientType:    "unknown",
				purchaseValue: 199,
			},
			0.00,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if gotDiscountRate := calculateDiscountRate(testCase.args.clientType, testCase.args.purchaseValue); gotDiscountRate != testCase.wantDiscountRate {
				t.Errorf("calculateDiscountRate() = %v, want %v", gotDiscountRate, testCase.wantDiscountRate)
			}
		})
	}
}
