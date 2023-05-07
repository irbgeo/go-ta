package ta

import (
	"github.com/shopspring/decimal"
)

// StDev - Standard Deviation.
func StDev(source Series, period int) Series {
	out := make(Series, 0, len(source))

	for i, v := range source {
		if i < period-1 {
			out = append(out, Value{
				Time:  v.Time,
				Value: decimal.Zero,
			})
			continue
		}
		v := stdev(source[i-period+1 : i+1])

		out = append(out, v)
	}

	return out
}

func stdev(source Series) Value {
	sma := sma(source)

	sum := decimal.Zero
	for _, v := range source {
		diff := v.Value.Sub(sma.Value)
		sum = sum.Add(diff.Pow(two))
	}

	length := decimal.NewFromInt(int64(len(source)))
	stdev := sqrt(sum.Div(length))
	return Value{
		Time:  source[len(source)-1].Time,
		Value: stdev,
	}
}
