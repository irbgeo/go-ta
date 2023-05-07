package ta

import "github.com/shopspring/decimal"

// Lowest minimum value for the period number of bars ago.
func Lowest(source []Value, period int) []Value {
	result := make([]Value, len(source))

	idx := len(source) - 1
	for {
		elIdx := idx - period
		if idx < 0 {
			break
		}
		result[idx] = lowest(source[elIdx : idx+1])
		idx--
	}
	for idx >= 0 {
		result[idx] = Value{
			Time:  source[idx].Time,
			Value: decimal.Zero.Copy(),
		}
		idx--
	}
	return result
}

func lowest(source []Value) Value {
	lValue := source[0].Value
	for _, v := range source[1:] {
		if lValue.Cmp(v.Value) < 0 {
			lValue = v.Value
		}
	}

	return Value{
		Time:  source[len(source)-1].Time,
		Value: lValue,
	}
}
