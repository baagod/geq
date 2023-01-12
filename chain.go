package ged

import (
	"fmt"
	"reflect"
	"strings"
)

type Eq map[string]any

// Where 返回带上 WHERE 前缀的 sql 字符串
//func (m And) Where() string {
//	if sql := m.SQL(); sql != "" {
//		return "WHERE " + sql
//	}
//	return ""
//}

func (e Eq) Cut() map[string]any {
	return cut(e)
}

func (e Eq) AND(a Eq) Cond {
	return Cond{list: []Eq{e}}.AND(a)
}

func (e Eq) OR(a Eq) Cond {
	return Cond{list: []Eq{e}}.OR(a)
}

func (e Eq) OrSQL() string {
	return toSQL(e, "OR")
}

func (e Eq) SQL() string {
	return toSQL(e, "AND")
}

func toSQL(m map[string]any, chain string) string {
	sql, op := "", "="

	for k, x := range m {
		if value, ok := x.(*Value); ok {
			if value.IsSkip() { // 排除该字段
				continue
			}

			if a, ok := value.value.(*or); ok {
				for _, v := range a.values {
					sql += fmt.Sprintf("%s %s %v OR ", k, v.Operate(), v.Out())
				}
				sql = strings.TrimRight(sql, " OR ") + chain
				continue
			} else {
				op = value.Operate()
			}
		}

		t := reflect.ValueOf(x)
		if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
			sql += fmt.Sprintf("%s %s %s ", k, inSQL(t), chain)
			continue
		}

		if _, ok := x.(string); ok {
			sql += fmt.Sprintf("%s %s '%v' %s ", k, op, x, chain)
		} else {
			sql += fmt.Sprintf("%s %s %v %s ", k, op, x, chain)
		}
	}

	return strings.TrimRight(sql, " "+chain+" ")
}

func inSQL(list reflect.Value) string {
	length, sql := list.Len(), ""

	for i := 0; i < length; i++ {
		x := list.Index(i).Interface() // 列表元素
		if v, ok := x.(*Value); ok && v.IsSkip() {
			continue
		}

		if _, ok := x.(string); ok {
			sql += fmt.Sprintf("'%v', ", x)
		} else {
			sql += fmt.Sprintf("%v, ", x)
		}
	}

	return fmt.Sprintf("IN (%s)", strings.TrimRight(sql, ", "))
}

// cut 返回过滤后的 map
func cut(m map[string]any) map[string]any {
	for k, v := range m {
		mv := reflect.ValueOf(v)                                      // map 中每个值的反射
		if mv.Kind() == reflect.Slice || mv.Kind() == reflect.Array { // map 的值为切片或数组
			for i := mv.Len() - 1; i >= 0; i-- { // 正序删除 mv 元素会出错
				v = mv.Index(i).Interface()
				if val, ok := v.(*Value); ok && val.IsSkip() { // 忽略值
					mv = reflect.AppendSlice(mv.Slice(0, i), mv.Slice(i+1, mv.Len())) // 移除第 i 个元素
				}
			}

			if mv.Len() == 0 {
				delete(m, k)
				continue
			}

			if val, ok := v.(*Value); ok {
				if val.isExpr = true; val.IsSkip() { // 排除该字段
					delete(m, k)
					continue
				}

				m[k] = val.Out()
				continue
			}

			m[k] = mv.Interface()
		}

		if val, ok := v.(*Value); ok {
			if val.isExpr = true; val.IsSkip() { // 排除该字段
				delete(m, k)
				continue
			}

			m[k] = val.Out()
		}
	}

	return m
}
