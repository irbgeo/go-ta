package ta

import "github.com/shopspring/decimal"

type Series []Value

func (s Series) Add(a Series) Series {
	if len(s) != len(a) {
		panic("mismatch series ")
	}

	out := make([]Value, len(s))
	for i := range s {
		out[i] = s[i].Add(a[i])
	}
	return out
}

func (s Series) Sub(a Series) Series {
	if len(s) != len(a) {
		panic("mismatch series ")
	}

	out := make([]Value, len(s))
	for i := range s {
		out[i] = s[i].Sub(a[i])
	}
	return out
}

func (s Series) Div(a Series) Series {
	if len(s) != len(a) {
		panic("mismatch series ")
	}

	out := make([]Value, len(s))
	for i := range s {
		out[i] = s[i].Div(a[i])
	}
	return out
}

func (s Series) MulConst(c decimal.Decimal) Series {
	out := make([]Value, len(s))
	for i, v := range s {
		out[i] = Value{
			Time:  v.Time,
			Value: v.Value.Mul(c),
		}
	}
	return out
}

func (s Series) DivConst(c decimal.Decimal) Series {
	out := make([]Value, len(s))
	for i, v := range s {
		out[i] = Value{
			Time:  v.Time,
			Value: v.Value.Div(c),
		}
	}
	return out
}

func (s Series) MaxConst(c decimal.Decimal) Series {
	out := make([]Value, len(s))
	for i, v := range s {
		out[i] = Value{
			Time: v.Time,
		}
		if v.Value.Cmp(c) > 0 {
			out[i].Value = v.Value.Copy()
			continue
		}
		out[i].Value = c.Copy()
	}
	return out
}

func (s Series) MinConst(c decimal.Decimal) Series {
	out := make([]Value, len(s))
	for i, v := range s {
		out[i] = Value{
			Time: v.Time,
		}
		if v.Value.Cmp(c) < 0 {
			out[i].Value = v.Value.Copy()
			continue
		}
		out[i].Value = c.Copy()
	}
	return out
}

func (s Series) Len() int {
	return len(s)
}
