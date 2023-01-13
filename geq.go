package ged

// Gt 大于 >
func Gt(v any) *Value {
	return &Value{value: v, operator: ">"}
}

// Ge 大于 ≥
func Ge(v any) *Value {
	return &Value{value: v, operator: ">="}
}

// Lt 小于 <
func Lt(v any) *Value {
	return &Value{value: v, operator: "<"}
}

// Le ≤
func Le(v any) *Value {
	return &Value{value: v, operator: "<="}
}

// Ne 不等于 != <>
func Ne(v any) *Value {
	return &Value{value: v, operator: "!="}
}

// Between 区间
func Between(first, second any) *Value {
	if _, ok := first.(*Value); !ok {
		first = &Value{value: first}
	}

	if _, ok := second.(*Value); !ok {
		second = &Value{value: second}
	}

	return &Value{
		value:    []*Value{first.(*Value), second.(*Value)},
		operator: "BETWEEN",
	}
}

// Or 或者
func Or(v ...any) *Value {
	var values []*Value
	for _, x := range v {
		if _, ok := x.(*Value); !ok {
			x = &Value{value: x}
		}
		values = append(values, x.(*Value))
	}
	return &Value{value: values, operator: "OR"}
}

// Expr 将值标识为表达式
// 表达式的值输出不带单引号(')
func Expr(v any) *Value {
	return &Value{value: v, expr: true}
}

// Skip 根据 b 忽略该值
func Skip(v any, b bool) *Value {
	return &Value{value: v, skip: b}
}

// SkipFunc 根据 fn 忽略该值
func SkipFunc[T any](v T, fn func(v T) bool) *Value {
	return &Value{value: v, skip: fn(v)}
}

// NonZero 非零的值，如果值为零值会被排除掉。
func NonZero(v any) *Value {
	return &Value{value: v, nonZero: true}
}
