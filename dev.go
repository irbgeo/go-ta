package ta

import (
	"github.com/shopspring/decimal"
)

// Dev returns of the difference between the series and its sma.
func Dev(source Series, period int) Series {
	out := make(Series, 0, len(source))

	for i, v := range source {
		if i < period-1 {
			out = append(out, Value{
				Time:  v.Time,
				Value: decimal.Zero,
			})
			continue
		}
		v := dev(source[i-period+1 : i+1])

		out = append(out, v)
	}

	return out
}

func dev(source Series) Value {
	sma := sma(source)
	sum := decimal.Zero

	for _, v := range source {
		diff := v.Value.Sub(sma.Value).Abs()
		sum = sum.Add(diff.Pow(two))
	}

	length := decimal.NewFromInt(int64(len(source)))
	dev := sqrt(sum.Div(length))
	return Value{
		Time:  source[len(source)-1].Time,
		Value: dev,
	}
}
