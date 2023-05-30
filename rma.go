package ta

import (
	"github.com/shopspring/decimal"
)

// RMA - Rolling Moving Average
// (rma[i-1]*(period-1)+src[i])/period
// https://download.esignal.com/products/workstation/help/charts/studies/rmi.htm
func RMA(src Series, period int) Series {
	if src.Len() < period {
		return nil
	}

	rma := make(Series, len(src))
	periodDec := decimal.NewFromInt(int64(period))

	sma := SMA(src, period)

	copy(rma, sma[:period])

	for i := period; i < src.Len(); i++ {
		rmaValue := (rma[i-1].Value.Mul(periodDec.Sub(one)).Add(src[i].Value)).Div(periodDec)
		rma[i] = Value{
			Time:  src[i].Time,
			Value: rmaValue,
		}

	}

	return rma
}
