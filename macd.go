package ta

// MACD - Moving Average Convergence/Divergence
func MACD(src Series, shortPeriod, longPeriod, signalPeriod int) (macd, signal, hist Series) {
	shortEMA := EMA(src, shortPeriod)
	longEMA := EMA(src, longPeriod)
	macd = shortEMA.Sub(longEMA)
	signal = EMA(macd, signalPeriod)
	hist = macd.Sub(signal)

	return
}
