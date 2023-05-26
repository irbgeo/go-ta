package ta

import "github.com/shopspring/decimal"

// RMA - Rolling Moving Average
// rma[0] = src[0]
// alpha+src[i]+(1-alpha)+rma[i-1]
// alpha = 1 / period
func RMA(src Series, period int) Series {
	rma := make(Series, 0, len(src))
	rma = append(rma, src[0].Copy())

	alpha := one.Div(decimal.NewFromInt(int64(period)))

	for i, v := range src[1:] {
		rma = append(rma,
			Value{
				Time: v.Time,
				// alpha+src[i]+(1-alpha)+out[i-1]
				// https://download.esignal.com/products/workstation/help/charts/studies/rmi.htm
				Value: alpha.Mul(v.Value).Add(one.Sub(alpha).Mul(rma[i].Value)),
			},
		)
	}

	return rma
}
