package ta

import (
	"github.com/shopspring/decimal"
)

// EMA - Exponential Moving Average
// EMA = alpha * source + (1 - alpha) * EMA[1]
// alpha = 2 / (length + 1)
// https://www.tradingview.com/pine-script-reference/v5/#fun_ta.ema
func EMA(src Series, period int) Series {
	if len(src) < period {
		return nil
	}

	alpha := decimal.NewFromFloat(2.0 / float64(period+1))
	result := make(Series, len(src))

	result[0] = src[0].Copy()

	for i, val := range src[1:] {
		ema := alpha.Mul(val.Value).Add(one.Sub(alpha).Mul(result[i].Value))

		result[i+1] = Value{Time: val.Time, Value: ema}
	}

	return result
}
