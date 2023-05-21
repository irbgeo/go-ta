package ta

// MACD - Moving Average Convergence/Divergence
func MACD(source Series, shortPeriod, longPeriod, signalPeriod int) (macd, signal, hist Series) {
	shortEMA := EMA(source, shortPeriod)
	longEMA := EMA(source, longPeriod)
	macd = shortEMA.Sub(longEMA)
	signal = EMA(macd, signalPeriod)
	hist = macd.Sub(signal)

	return
}
