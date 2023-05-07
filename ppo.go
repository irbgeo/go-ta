package ta

// PPO - Percentage Price Oscillator
// ppo = 100*(shortEMA - longEMA)/longEMA
func PPO(source Series, longPeriod, shortPeriod int) Series {
	longEMA := EMA(source, longPeriod)
	shortEMA := EMA(source, shortPeriod)

	return shortEMA.Sub(longEMA).Div(longEMA).MulConst(thousand)
}
