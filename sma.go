package ta

import "github.com/shopspring/decimal"

// SMA - Simple Moving Average
func SMA(source Series, period int) Series {
	out := make(Series, 0, len(source))

	sum := decimal.Zero
	for i, v := range source {
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
		sum = sum.Sub(source[i-period+1].Value)
	}

	return out
}

func sma(source Series) Value {
	sum := decimal.Zero
	for _, v := range source {
		sum = sum.Add(v.Value)
	}

	return Value{
		Time:  source[len(source)-1].Time,
		Value: sum.Div(decimal.NewFromInt(int64(len(source)))),
	}
}
