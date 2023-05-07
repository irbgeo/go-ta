package ta

import "github.com/shopspring/decimal"

const defaultPeriod = 1

// Change compares the current source value with its period value a few bars ago and returns the difference.
// default period = 1
func Change(source Series, period ...int) Series {
	p := defaultPeriod
	if len(period) != 0 {
		p = period[0]
	}

	result := make(Series, len(source))

	idx := len(source) - 1
	for {
		elIdx := idx - p
		if elIdx < 0 {
			break
		}
		result[idx] = source[idx].Sub(source[elIdx])
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
