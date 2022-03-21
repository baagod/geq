package eq

func Eq(v interface{}) Value { // 大于 >
	return Value{value: v}.gt(true)
}

func Gt(v interface{}) Value { // 大于 >
	return Value{value: v}.gt(true)
}

func Ge(v interface{}) Value { // 大于 ≥
	return Value{value: v}.ge(true)
}

func Lt(v interface{}) Value { // 小于 <
	return Value{value: v}.lt(true)
}

func Le(v interface{}) Value { // ≤
	return Value{value: v}.le(true)
}

func Ne(v interface{}) Value { // 不等于 ≠
	return Value{value: v}.ne(true)
}

// ---

func Expr(v interface{}) Value { // 是否表达式，表达式的值会原样输出，不带单引号[‘’]
	return Value{value: v}.Expr(true)
}

func Not[T any](v T, operator func(v T) bool) Value { // 是否排除该值
	return Value{value: v}.Not(operator(v))
}

func NonZero(v interface{}) Value { // 非零值，零值会被过滤掉
	return Value{value: v}.NonZero()
}

// ----

func Where(eq Builder) Container { // 生成条件容器
	if _, ok := eq.(And); ok {
		return Container{list: []Builder{eq}, operators: []string{"AND"}}
	}
	return Container{list: []Builder{eq}, operators: []string{"OR"}}
}
