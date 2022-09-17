package eq

import "fmt"
import "reflect"

type Value struct {
	value     any
	isEq      bool // 等于 =
	isGt      bool // 大于 	   >
	isGe      bool // 大于等于 ≥
	isLt      bool // 小于	   <
	isLe      bool // 小于等于 ≤
	isNe      bool // 不等于   ≠
	isBetween bool // 区间、范围 BETWEEN
	isExpr    bool // 是否表达式，表达式的值会原样输出，不带单引号('')。
	isOut     bool // 是否排除该值
	isNoZero  bool // 非零值，零值会被排除。
}

// Out 根据给定的条件排除值
func (v Value) Out(b bool) Value {
	v.isOut = b
	return v
}

// Expr 标识值为表达式，表达式的值会原样输出，不带单引号(”)。
func (v Value) Expr() Value {
	v.isExpr = true
	return v
}

// NoZero 标识为非零值，零值会被排除。
func (v Value) NoZero() Value {
	v.isNoZero = true
	return v
}

// eq 等于条件
func (v Value) eq() Value {
	v.isEq = true
	return v
}

// gt > 大于
func (v Value) gt() Value {
	v.isGt = true
	return v
}

// ge >= 大于等于
func (v Value) ge() Value {
	v.isGe = true
	return v
}

// lt < 小于
func (v Value) lt() Value {
	v.isLt = true
	return v
}

// le <= 小于等于
func (v Value) le() Value {
	v.isLe = true
	return v
}

// ne <>, != 不等于
func (v Value) ne() Value {
	v.isNe = true
	return v
}

// ne <>, != 不等于
func (v Value) between() Value {
	v.isBetween = true
	return v
}

func (v Value) String() string {
	if v.isExpr {
		if _, ok := v.value.(string); ok {
			return fmt.Sprintf("'%v'", v.value)
		}
	} else if v.isBetween {
		slice := reflect.ValueOf(v.value).Slice(0, 2)
		first := slice.Index(0).Interface()
		second := slice.Index(1).Interface()

		if _, ok := first.(string); ok {
			first = fmt.Sprintf("'%s'", first)
		}

		if _, ok := second.(string); ok {
			second = fmt.Sprintf("'%s'", second)
		}

		return fmt.Sprintf("%v AND %v", first, second)
	}

	return fmt.Sprint(v.value)
}
