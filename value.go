package eq

import "fmt"
import "reflect"

type Value struct {
    value     any
    isEq      bool // 等于 =
    isGt      bool // 大于 >
    isGe      bool // 大于等于 ≥
    isLt      bool // 小于 <
    isLe      bool // 小于等于 ≤
    isNe      bool // 不等于 ≠
    isBetween bool // 区间、范围 BETWEEN
    isExpr    bool // 是否表达式，标识为表达式的输出不带单引号('')。
    isSkip    bool // 是否忽略该值
    isNoZero  bool // 是否忽略零值
}

// Skip 根据给定的条件忽略该值
func (v *Value) Skip(b bool) *Value {
    v.isSkip = b
    return v
}

// Expr 标识为表达式的输出不带单引号(”)
func (v *Value) Expr() *Value {
    v.isExpr = true
    return v
}

// NoZero 排除零值
func (v *Value) NoZero() *Value {
    v.isNoZero = true
    return v
}

func (v *Value) String() string {
    return fmt.Sprint(v.val())
}

// 私有方法

// val 返回具体包装在 Value 中的值
func (v *Value) val() any {
    if v.isBetween {
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

    if _, ok := v.value.(string); ok && !v.isExpr {
        return fmt.Sprintf("'%s'", v.value)
    }

    return v.value
}

// ignore 如果该值是被忽略的，则返回 true，否则为 false。
func (v *Value) ignore() bool {
    return v.isSkip || v.isNoZero && reflect.ValueOf(v.value).IsZero()
}
