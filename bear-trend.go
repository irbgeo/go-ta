package ta

func BearTrend(high, close, low Series) Series {
	high = high[:len(high)-1]
	close = close[:len(close)-1]
	low = low[:len(low)-1]

	bearTrend := Highest(high, 50).Sub(close).Div(ATR(high, close, low, 5))

	return bearTrend
}
