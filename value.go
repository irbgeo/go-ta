package ta

import (
	"github.com/shopspring/decimal"
)

type Value struct {
	Time  int64
	Value decimal.Decimal
}

// c = a+b
func (a Value) Add(b Value) (c Value) {
	c = Value{
		Time:  a.Time,
		Value: a.Value.Add(b.Value),
	}

	return
}

// c = a-b
func (a Value) Sub(b Value) (c Value) {
	c = Value{
		Time:  a.Time,
		Value: a.Value.Sub(b.Value),
	}

	return
}

// c = b*a
func (a Value) Mul(b Value) (c Value) {
	c = Value{
		Time:  a.Time,
		Value: a.Value.Mul(b.Value),
	}

	return
}

// c = a/b
func (a Value) Div(b Value) (c Value) {
	c = Value{
		Time:  a.Time,
		Value: a.Value.Div(b.Value),
	}

	return
}

// c = abs(a)
func (a Value) Abs() (c Value) {
	c = Value{
		Time:  a.Time,
		Value: a.Value.Abs(),
	}

	return
}

// c = abs(a)
func (a Value) Sqrt() (c Value) {
	c = Value{
		Time:  a.Time,
		Value: sqrt(a.Value),
	}

	return
}

// b = c*a
func (a Value) MulConst(c decimal.Decimal) (b Value) {
	b = Value{
		Time:  a.Time,
		Value: a.Value.Mul(c),
	}

	return
}

// b = a/c
func (a Value) DivConst(c decimal.Decimal) (b Value) {
	b = Value{
		Time:  a.Time,
		Value: a.Value.Div(c),
	}

	return
}

func (v Value) Copy() Value {
	return Value{
		Time:  v.Time,
		Value: v.Value.Copy(),
	}
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (a Value) Cmp(b Value) int {
	return a.Value.Cmp(b.Value)
}
