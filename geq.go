package ged

// V 等于 =
func V(v any) *Value {
	return &Value{value: v, isEq: true}
}

// Gt 大于 >
func Gt(v any) *Value {
	return &Value{value: v, isGt: true}
}

// Ge 大于 ≥
func Ge(v any) *Value {
	return &Value{value: v, isGe: true}
}

// Lt 小于 <
func Lt(v any) *Value {
	return &Value{value: v, isLt: true}
}

// Le ≤
func Le(v any) *Value {
	return &Value{value: v, isLe: true}
}

// Ne 不等于 ≠
func Ne(v any) *Value {
	return &Value{value: v, isNe: true}
}

func Or(v ...any) *Value {
	var values []*Value
	for _, x := range v {
		if _, ok := x.(*Value); !ok {
			x = &Value{value: x}
		}
		values = append(values, x.(*Value))
	}
	return &Value{value: &or{values}, isOr: true}
}

// Between 区间
func Between(first, second any) *Value {
	Or(1, 10.2, "100")
	var value between
	if _, ok := first.(*Value); !ok {
		value.first = &Value{value: first}
	}
	if _, ok := second.(*Value); !ok {
		value.second = &Value{value: second}
	}
	return &Value{value: &value, isBetween: true}
}

// Expr 将值标识为表达式
// 表达式的值输出不带单引号(')
func Expr(v any) *Value {
	return &Value{value: v, isExpr: true}
}

// Skip 根据 b 忽略该值
func Skip(v any, b bool) *Value {
	return &Value{value: v, isSkip: b}
}

// SkipFunc 根据 fn 忽略该值
func SkipFunc[T any](v T, fn func(v T) bool) *Value {
	return &Value{value: v, isSkip: fn(v)}
}

// NoZero 非零的值，如果值为零值会被排除掉。
func NoZero(v any) *Value {
	return &Value{value: v, isNoZero: true}
}

// ----

// Where 生成条件集合
//func Where(m Builder) Container {
//	if _, ok := m.(And); ok {
//		return Container{list: []Builder{m}, operators: []string{"AND"}}
//	}
//	return Container{list: []Builder{m}, operators: []string{"OR"}}
//}
