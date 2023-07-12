package ta

import "github.com/shopspring/decimal"

// Highest maximum value for the period number of bars ago.
func Highest(src Series, period int) Series {
	result := make(Series, len(src))

	eIdx := len(src) - 1
	for {
		bIdx := eIdx - period + 1
		if bIdx < 0 {
			break
		}
		result[eIdx] = highest(src[bIdx : eIdx+1])
		eIdx--
	}
	for eIdx >= 0 {
		result[eIdx] = Value{
			Time:  src[eIdx].Time,
			Value: decimal.Zero.Copy(),
		}
		eIdx--
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
