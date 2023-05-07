package ta

import (
	"github.com/shopspring/decimal"
)

// research
func sqrt(x decimal.Decimal) decimal.Decimal {
	f := x.BigFloat()
	v, _ := f.Sqrt(f).Float64()

	return decimal.NewFromFloat(v)
}

func max(arr ...Value) Value {
	if len(arr) == 0 {
		return Value{}
	}

	max := arr[0]
	for _, v := range arr[1:] {
		if max.Cmp(v) < 0 {
			max = v
		}
	}

	return max
}
