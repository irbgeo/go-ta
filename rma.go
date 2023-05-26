package ta

import "github.com/shopspring/decimal"

// RMA - Rolling Moving Average
func RMA(src Series, period int) Series {
	out := make(Series, 0, len(src))
	out = append(out, src[0].Copy())

	alpha := one.Div(decimal.NewFromInt(int64(period)))

	for i, v := range src[1:] {
		out = append(out,
			Value{
				Time:  v.Time,
				Value: alpha.Mul(v.Value).Add(one.Sub(alpha).Mul(out[i].Value)),
			},
		)
	}

	return out
}
