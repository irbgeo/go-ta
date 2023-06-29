package ta

import "github.com/shopspring/decimal"

// Highest maximum value for the period number of bars ago.
func Highest(src Series, period int) Series {
	result := make(Series, len(src))

	idx := len(src) - 1
	for {
		bIdx := idx - period
		if bIdx < 0 {
			break
		}
		result[idx] = highest(src[bIdx : idx+1])
		idx--
	}
	for idx >= 0 {
		result[idx] = Value{
			Time:  src[idx].Time,
			Value: decimal.Zero.Copy(),
		}
		idx--
	}

	return result
}

func highest(src Series) Value {
	lValue := src[0].Value
	for _, v := range src[1:] {
		if lValue.Cmp(v.Value) < 0 {
			lValue = v.Value
		}
	}

	return Value{
		Time:  src[len(src)-1].Time,
		Value: lValue,
	}
}
