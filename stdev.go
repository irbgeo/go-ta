package ta

import (
	"github.com/shopspring/decimal"
)

// StDev - Standard Deviation.
func StDev(src Series, period int) Series {
	out := make(Series, 0, len(src))

	for i, v := range src {
		if i < period-1 {
			out = append(out, Value{
				Time:  v.Time,
				Value: decimal.Zero,
			})
			continue
		}
		v := stdev(src[i-period+1 : i+1])

		out = append(out, v)
	}

	return out
}

func stdev(src Series) Value {
	sma := sma(src)

	sum := decimal.Zero
	for _, v := range src {
		diff := v.Value.Sub(sma.Value)
		sum = sum.Add(diff.Pow(two))
	}

	length := decimal.NewFromInt(int64(len(src)))
	stdev := sqrt(sum.Div(length))
	return Value{
		Time:  src[len(src)-1].Time,
		Value: stdev,
	}
}
