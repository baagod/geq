package eq

// Eq 等于 >
func Eq(v any) Value {
	return Value{value: v}.eq()
}

// Gt 大于 >
func Gt(v any) Value {
	return Value{value: v}.gt()
}

// Ge 大于 ≥
func Ge(v any) Value {
	return Value{value: v}.ge()
}

// Lt 小于 <
func Lt(v any) Value {
	return Value{value: v}.lt()
}

// Le ≤
func Le(v any) Value {
	return Value{value: v}.le()
}

// Ne 不等于 ≠
func Ne(v any) Value {
	return Value{value: v}.ne()
}

func Between(first any, second any) Value {
	return Value{value: []any{first, second}}.between()
}

// ---

func Expr(v any) Value { // 是否表达式，表达式的值会原样输出，不带单引号[‘’]
	return Value{value: v}.Expr()
}

func Out[T any](v T, operator func(v T) bool) Value { // 是否排除该值
	return Value{value: v}.Out(operator(v))
}

func NoZero(v any) Value { // 非零值，零值会被过滤掉
	return Value{value: v}.NoZero()
}

// ----

func Where(eq Builder) Container { // 生成条件容器
	if _, ok := eq.(And); ok {
		return Container{list: []Builder{eq}, operators: []string{"AND"}}
	}
	return Container{list: []Builder{eq}, operators: []string{"OR"}}
}
