package ta

import "github.com/shopspring/decimal"

var scaledIndex = decimal.NewFromFloat(0.015)

// CCI - Commodity Channel Index
// cci = (src-sma)/(0.015*dev)
func CCI(src Series, period int) Series {
	sma := SMA(src, period)
	dev := Dev(src, period)

	return src.Sub(sma).Div(dev.DivConst(scaledIndex))
}
