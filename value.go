package ged

import (
	"fmt"
	"reflect"
)

type between struct {
	first  *Value
	second *Value
}

type or struct {
	values []*Value
}

type Value struct {
	value     any
	isEq      bool // 等于 =
	isGt      bool // 大于 >
	isGe      bool // 大于等于 ≥
	isLt      bool // 小于 <
	isLe      bool // 小于等于 ≤
	isNe      bool // 不等于 ≠
	isOr      bool // 或者 or
	isBetween bool // 区间 BETWEEN
	isExpr    bool // 是否表达式
	isSkip    bool // 是否跳过值
	isNoZero  bool // 是否非零值
}

func (v *Value) Expr() *Value {
	v.isExpr = true
	return v
}

func (v *Value) Skip() *Value {
	v.isSkip = true
	return v
}

func (v *Value) NoZero() *Value {
	v.isNoZero = true
	return v
}

func (v *Value) String() string {
	return fmt.Sprint(v.Out())
}

func (v *Value) Operate() string {
	if v.isGt {
		return ">"
	} else if v.isGe {
		return ">="
	} else if v.isLt {
		return "<"
	} else if v.isLe {
		return "<="
	} else if v.isNe {
		return "!="
	} else if v.isBetween {
		return "BETWEEN"
	}
	return "="
}

// IsSkip 返回值是否被跳过
func (v *Value) IsSkip() bool {
	return v.isSkip ||
		v.isNoZero && reflect.ValueOf(v.value).IsZero()
}

// Out 返回输出的 value 值
func (v *Value) Out() any {
	if b, ok := v.value.(*between); ok {
		return fmt.Sprintf("%v AND %v", b.first.Out(), b.second.Out())
	}

	if _, ok := v.value.(string); ok && !v.isExpr {
		return fmt.Sprintf("'%s'", v.value)
	}

	return v.value
}
