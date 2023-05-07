package ta

import "github.com/shopspring/decimal"

var scaledIndex = decimal.NewFromFloat(0.015)

// CCI - Commodity Channel Index
// cci = (source-sma)/(0.015*dev)
func CCI(source Series, period int) Series {
	sma := SMA(source, period)
	dev := Dev(source, period)

	return source.Sub(sma).Div(dev.DivConst(scaledIndex))
}
