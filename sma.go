package ta

import "github.com/shopspring/decimal"

// SMA - Simple Moving Average
func SMA(src Series, period int) Series {
	out := make(Series, 0, len(src))

	sum := decimal.Zero
	for i, v := range src {
		sum = sum.Add(v.Value)
		if i < period-1 {
			out = append(out, Value{
				Time:  v.Time,
				Value: decimal.Zero,
			})
			continue
		}

		sma := sum.Div(decimal.NewFromInt(int64(period)))
		out = append(
			out,
			Value{
				Time:  v.Time,
				Value: sma,
			},
		)
		sum = sum.Sub(src[i-period+1].Value)
	}

	return out
}

func sma(src Series) Value {
	sum := decimal.Zero
	for _, v := range src {
		sum = sum.Add(v.Value)
	}

	return Value{
		Time:  src[len(src)-1].Time,
		Value: sum.Div(decimal.NewFromInt(int64(len(src)))),
	}
}
