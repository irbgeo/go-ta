package ta

import (
	"github.com/shopspring/decimal"
)

var scaledIndex = decimal.New(15, -3)

// CCI - Commodity Channel Index
// cci = (src-sma)/(0.015*dev)
// https://www.tradingview.com/pine-script-reference/v5/#fun_ta.cci
func CCI(high, low, close Series, period int) Series {
	if len(high) != len(low) || len(low) != len(close) || len(high) != len(close) {
		panic("length mismatch")
	}

	hlc3 := make(Series, 0, high.Len())
	for i := 0; i < len(high); i++ {
		el := high[i].Add(low[i]).Add(close[i]).DivConst(three)

		hlc3 = append(hlc3, el)
	}

	sma := SMA(hlc3, period)
	dev := Dev(hlc3, period)

	return hlc3.Sub(sma).Div(dev.MultiConst(scaledIndex))
}
