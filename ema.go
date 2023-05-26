package ta

import "github.com/shopspring/decimal"

// EMA - Exponential Moving Average
func EMA(src Series, period int) Series {
	if len(src) < period {
		return nil
	}

	ema := src[0].Value
	multiplier := decimal.NewFromFloat(2.0 / float64(period+1))
	result := make(Series, len(src))

	for i, val := range src[1:] {
		ema = val.Value.Sub(ema).Mul(multiplier).Add(ema)

		result[i] = Value{Time: val.Time, Value: ema}
	}

	return result
}
