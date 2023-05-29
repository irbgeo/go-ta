package ta

import (
	"github.com/shopspring/decimal"
)

// RMA - Rolling Moving Average
// (rma[i-1]*(period-1)+src[i])/period
// https://download.esignal.com/products/workstation/help/charts/studies/rmi.htm
func RMA(src Series, period int) Series {
	rma := make(Series, 0, len(src))
	periodDec := decimal.NewFromInt(int64(period))

	rma = append(rma,
		Value{
			Time:  src[0].Time,
			Value: src[0].Value.Div(periodDec),
		},
	)

	for i, v := range src[1:] {
		rmaValue := (rma[i].Value.Mul(periodDec.Sub(one)).Add(v.Value)).Div(periodDec)
		rma = append(rma,
			Value{
				Time:  v.Time,
				Value: rmaValue,
			},
		)
	}

	return rma
}
