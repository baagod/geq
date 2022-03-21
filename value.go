package eq

import "fmt"

type Value struct {
	value     interface{}
	isGt      bool // 大于 	   >
	isGe      bool // 大于等于 ≥
	isLt      bool // 小于	   <
	isLe      bool // 小于等于 ≤
	isNe      bool // 不等于   ≠
	isExpr    bool // 是否表达式，表达式的值会原样输出，不带单引号[‘’]。
	isNot     bool // 是否排除该值
	isNonZero bool // 非零值，零值会被排除。
}

func (v Value) Not(b bool) Value {
	v.isNot = b
	return v
}

func (v Value) Expr(b bool) Value {
	v.isExpr = b
	return v
}

func (v Value) NonZero() Value {
	v.isNonZero = true
	return v
}

func (v Value) gt(b bool) Value {
	v.isGt = b
	return v
}

func (v Value) ge(b bool) Value {
	v.isGe = b
	return v
}

func (v Value) lt(b bool) Value {
	v.isLt = b
	return v
}

func (v Value) le(b bool) Value {
	v.isLe = b
	return v
}

func (v Value) ne(b bool) Value {
	v.isNe = b
	return v
}

func (v Value) String() string {
	return fmt.Sprint(v.value)
}
