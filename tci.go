package ta

import "github.com/shopspring/decimal"

// TCI - Triple Cross Indicator
func TCI(source Series,
	shortMACDPeriod, longMACDPeriod int) Series {

	macd, _, _ := MACD(source, shortMACDPeriod, longMACDPeriod, 0)
	hibRSI := RSI(source, 13).Sub(RSI(source, 65))
	hibPPO := PPO(source, 12, 26).Sub(PPO(source, 55, 80))

	tci := make(Series, len(source))
	for i := range tci {
		tciValue := decimal.Zero.Copy()

		if macd[i].Value.Cmp(decimal.Zero) > 0 {
			tciValue.Add(one)
		} else {
			tciValue.Sub(one)
		}

		if hibRSI[i].Value.Cmp(decimal.Zero) > 0 {
			tciValue.Add(one)
		} else {
			tciValue.Sub(one)
		}

		if hibPPO[i].Value.Cmp(decimal.Zero) > 0 {
			tciValue.Add(one)
		} else {
			tciValue.Sub(one)
		}
	}

	return tci
}
