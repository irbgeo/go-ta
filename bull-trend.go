package ta

func BullTrend(high, close, low Series) Series {
	high = high[:high.Len()-1]
	close = close[:close.Len()-1]
	low = low[:low.Len()-1]

	bullTrend := close.Sub(Highest(high, 50)).Div(ATR(high, close, low, 5))

	return bullTrend
}
