package ged

import (
	"fmt"
	"reflect"
	"strings"
)

type Eq map[string]any

func (e Eq) AND(a Eq) Cond {
	return Cond{list: []Eq{e}}.AND(a)
}

func (e Eq) OR(a Eq) Cond {
	return Cond{list: []Eq{e}}.OR(a)
}

func (e Eq) SQL() string {
	return toSQL(e, "AND")
}

func (e Eq) OrSQL() string {
	return toSQL(e, "OR")
}

// WhereSQL 返回带有 WHERE 前缀的 sql 字符串
func (e Eq) WhereSQL() string {
	if sql := e.SQL(); sql != "" {
		return "WHERE " + sql
	}
	return ""
}

func toSQL(m map[string]any, chain string) string {
	var sql string
	for k, x := range m {
		if val, ok := x.(*Value); ok {
			if val.IsSkip() { // 排除该字段
				continue
			}
			if val.operator == "OR" {
				for _, v := range val.value.([]*Value) {
					sql += fmt.Sprintf("%s %s %v OR ", k, operator(v), value(v))
				}
				sql = strings.TrimRight(sql, " OR ") + chain
				continue
			}
		}
		sql += fmt.Sprintf("%s %s %v %s ", k, operator(x), value(x), chain)
	}
	return strings.TrimRight(sql, " "+chain+" ")
}

// Clip 返回修剪后的 map
func (e Eq) Clip() map[string]any {
	m := map[string]any{}

	for k, v := range e {
		if val, ok := v.(*Value); ok {
			if val.IsSkip() { // 排除该字段
				continue
			}
		}

		t := reflect.ValueOf(v)                                     // map 中每个值的反射
		if t.Kind() == reflect.Slice || t.Kind() == reflect.Array { // map 的值为切片或数组
			for i := t.Len() - 1; i >= 0; i-- { // 正序删除元素会出错
				v = t.Index(i).Interface()
				if val, ok := v.(*Value); ok && val.IsSkip() { // 忽略值
					t = reflect.AppendSlice(t.Slice(0, i), t.Slice(i+1, t.Len())) // 移除第 i 个元素
				}
			}

			if t.Len() > 0 {
				m[k] = t.Interface()
			}

			continue
		}

		m[k] = value(v)
	}

	return m
}
