package ta

import "github.com/shopspring/decimal"

// RMA - Rolling Moving Average
func RMA(source Series, period int) Series{
	if len(source) < period {
		return nil
	}

	out := make(Series, 0, len(source))

	sum := decimal.Zero
	for _, v := range source[:period] {
		out = append(out, Value{
			Time:  v.Time,
			Value: decimal.Zero,
		})
		sum = sum.Add(v.Value)
	}

	p := decimal.NewFromInt(int64(period))
	out[period-1] = Value{
		Time:  source[period-1].Time,
		Value: sum.Div(p),
	}

	k := two.Div(p.Add(one))

	for i, v := range source[period-1:] {
		i += period - 1

		prevRMA := out[i-1].Value
		curRMA := (v.Value.Sub(prevRMA)).Mul(k).Add(prevRMA)
		out = append(out, Value{
			Time:  v.Time,
			Value: curRMA,
		})
	}

	return out
}
