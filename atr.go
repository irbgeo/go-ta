package ta

// ATR - Average True Range
// returns the RMA of true range
// True range is max(high - low, abs(high - prev close), abs(low - prev close)).
func ATR(high, close, low []Value, period int) []Value {
	if len(high) != len(close) || len(high) != len(low) || len(low) != len(close) {
		return nil
	}

	trueRange := make([]Value, len(high))
	for i := len(high) - 1; i > 0; i-- {
		trueRange[i] = max(high[i].Sub(low[i]), high[i].Sub(close[i-1]).Abs(), low[i].Sub(close[i-1]).Abs())
	}

	return RMA(trueRange, period)
}
