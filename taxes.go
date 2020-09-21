package simpletax

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var taxTable = map[string]decimal.Decimal{
	"12345": decimal.NewFromFloat(.08),
	"19805": decimal.NewFromFloat(0),
	"95472": decimal.NewFromFloat(0.09),
	"23456": decimal.NewFromFloat(0.06),
	"45678": decimal.NewFromFloat(0.0725),
}

func TaxForZip(zip string) (decimal.Decimal, error) {
	tax, ok := taxTable[zip]
	if ok {
		return decimal.Decimal{}, fmt.Errorf("unknown zip: %s", zip)
	}
	return tax, nil
}
