package ged

import (
	"fmt"
	"reflect"
	"strings"
)

type Value struct {
	value    any
	expr     bool
	skip     bool
	nonZero  bool
	operator string // 运算符
}

// IsSkip 返回值是否被跳过
func (v Value) IsSkip() bool {
	return v.skip ||
		v.nonZero && reflect.ValueOf(v.value).IsZero()
}

func (v Value) String() string {
	return fmt.Sprint(value(&v))
}

func operator(v any) (operator string) {
	if val, ok := v.(*Value); ok {
		operator = val.operator
	}
	if operator == "" {
		return "="
	}
	return
}

func value(v any) any {
	if val, ok := v.(*Value); ok {
		if x, ok := val.value.([]*Value); ok {
			return fmt.Sprintf("%s AND %s", x[0], x[1])
		}

		if s, ok := val.value.(string); ok && !val.expr {
			return fmt.Sprintf("'%s'", s)
		}

		return val.value
	}

	if _, ok := v.(string); ok {
		return fmt.Sprintf("'%s'", v)
	}

	t := reflect.ValueOf(v)
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		length, sql := t.Len(), ""
		for i := 0; i < length; i++ {
			x := t.Index(i).Interface() // 列表元素
			if v, ok := x.(*Value); ok && v.IsSkip() {
				continue
			}
			sql += fmt.Sprintf("%v, ", value(x))
		}
		return fmt.Sprintf("IN (%s)", strings.TrimRight(sql, ", "))
	}

	return v
}
