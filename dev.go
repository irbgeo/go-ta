package ta

import (
	"github.com/shopspring/decimal"
)

// Dev returns of the difference between the series and its sma.
// https://www.tradingview.com/pine-script-reference/v5/#fun_ta.dev
func Dev(src Series, period int) Series {
	out := make(Series, 0, len(src))

	for i, v := range src {
		if i < period-1 {
			out = append(out, Value{
				Time:  v.Time,
				Value: decimal.Zero,
			})
			continue
		}
		v := dev(src[i-period+1 : i+1])

		out = append(out, v)
	}

	return out
}

func dev(src Series) Value {
	sma := sma(src)
	sum := decimal.Zero

	for _, v := range src {
		diff := v.Value.Sub(sma.Value).Abs()
		sum = sum.Add(diff.Pow(two))
	}

	length := decimal.NewFromInt(int64(len(src)))
	dev := sqrt(sum.Div(length))
	return Value{
		Time:  src[len(src)-1].Time,
		Value: dev,
	}
}
