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
    isOut     bool // 是否排除该值
    isNoZero  bool // 是否排除零值
}

// Out 根据给定的条件排除值
func (v *Value) Out(b bool) *Value {
    v.isOut = b
    return v
}

// Expr 标识为表达式的输出不带单引号('')
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

    if v.isExpr {
        if _, ok := (v.value).(string); ok {
            return fmt.Sprintf("'%s'", v.value)
        }
    }

    if _, ok := (v.value).(string); ok {
        return fmt.Sprintf("'%s'", v.value)
    }

    return fmt.Sprint(v.value)
}