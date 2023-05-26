package ta

// PPO - Percentage Price Oscillator
// ppo = 100*(shortEMA - longEMA)/longEMA
func PPO(src Series, longPeriod, shortPeriod int) Series {
	longEMA := EMA(src, longPeriod)
	shortEMA := EMA(src, shortPeriod)

	return shortEMA.Sub(longEMA).Div(longEMA).MulConst(thousand)
}
