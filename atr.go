package ta

// ATR - Average True Range
// returns the RMA of true range
// True range is max(high - low, abs(high - prev close), abs(low - prev close)).
// https://www.tradingview.com/pine-script-reference/v5/#fun_ta.atr
func ATR(high, close, low Series, period int) Series {
	if len(high) != len(close) || len(high) != len(low) || len(low) != len(close) {
		return nil
	}

	trueRange := make(Series, len(high))
	for i := len(high) - 1; i > 0; i-- {
		trueRange[i] = max(high[i].Sub(low[i]), high[i].Sub(close[i-1]).Abs(), low[i].Sub(close[i-1]).Abs())
	}

	return RMA(trueRange, period)
}
