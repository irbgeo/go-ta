package ta

import "github.com/shopspring/decimal"

// RSI - Relative Strength Index
func RSI(src Series, period int) Series {
	change := Change(src)

	u := change.MaxConst(decimal.Zero)
	d := change.MultiConst(negOne).MaxConst(decimal.Zero)

	rs := RMA(u, period).Div(RMA(d, period))

	rsi := make(Series, 0, len(src))
	for _, v := range rs {
		// rsi = 100 - 100 / (1 + v)
		rsiValue := v.Value.Add(one)
		rsiValue = thousand.Div(rsiValue)
		rsiValue = thousand.Sub(rsiValue)

		rsi = append(
			rsi,
			Value{
				Time:  v.Time,
				Value: rsiValue,
			},
		)
	}
	return rsi
}
