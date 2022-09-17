package eq

// Eq 等于 >
func Eq(v any) *Value {
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

// Between 区间
func Between(first, second any) *Value {
    return &Value{value: []any{first, second}, isBetween: true}
}

// ----

// Expr 标识为表达式的输出不带单引号('')
func Expr(v any) Value {
    return Value{value: v, isExpr: true}
}

// Out 根据给定的条件排除值
func Out[T any](v T, operator func(v T) bool) *Value {
    return (&Value{value: v}).Out(operator(v))
}

// NoZero 排除零值
func NoZero(v any) *Value {
    return (&Value{value: v}).NoZero()
}

// ----

// Where 生成条件集合
func Where(m Builder) Container {
    if _, ok := m.(And); ok {
        return Container{list: []Builder{m}, operators: []string{"AND"}}
    }
    return Container{list: []Builder{m}, operators: []string{"OR"}}
}
