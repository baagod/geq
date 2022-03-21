package eq

import (
	"fmt"
	"reflect"
	"strings"
)

type Builder interface {
	ToSql() string
	Cut() Map
}

type Map map[string]interface{}

// ToSql 返回 sql 字符串
func (b Map) ToSql() string { return "" }

// Cut 返回过滤后的 And 对象
func (b Map) Cut() Map {
	for k, v := range b {
		mvv := reflect.ValueOf(v)                                       // map 中每个值的反射类型
		if mvv.Kind() == reflect.Slice || mvv.Kind() == reflect.Array { // map 的值为切片或数组类型
			for i := mvv.Len() - 1; i >= 0; i-- { // 正序删除 mvv 元素会出错
				v = mvv.Index(i).Interface()
				if val, ok := v.(Value); ok {
					if val.isNot || val.isNonZero && reflect.ValueOf(val.value).IsZero() { // 排除该元素
						mvv = reflect.AppendSlice(mvv.Slice(0, i), mvv.Slice(i+1, mvv.Len())) // 移除第 i 个元素
					}
				}
			}

			if mvv.Len() == 0 {
				delete(b, k)
				continue
			}

			b[k] = mvv.Interface()
		}

		if val, ok := v.(Value); ok {
			if val.isNot || val.isNonZero && reflect.ValueOf(val.value).IsZero() { // 排除该字段
				delete(b, k)
			}
		}
	}

	return b
}

// ---- 等于 AND ----

type And Map

func (m And) ToSql() string { return toSql(m, "AND") }

func (m And) Cut() Map { return Map(m).Cut() }

// ---- 或者 OR ----

type Or Map

func (m Or) ToSql() string { return toSql(m, "OR") }

func (m Or) Cut() Map { return Map(m).Cut() }

// ---- 私有方法 ----

func toSql(eq map[string]interface{}, chain string) string {
	var sql string
	operator := "="

	for k, v := range eq {
		mvv := reflect.ValueOf(v)
		_, isMark := v.(string) // 是否给值[v]加上单引号

		if mvv.Kind() == reflect.Slice || mvv.Kind() == reflect.Array {
			listSql := toInSql(mvv)
			sql += fmt.Sprintf("%s %s %s ", k, listSql, chain)
			continue
		}

		if val, ok := v.(Value); ok {
			vv := reflect.ValueOf(val.value)
			if val.isNot || val.isNonZero && vv.IsZero() { // 排除该字段
				continue
			}

			v = val.value // 改变 v 的值以追加到 sql
			isMark = vv.Type().Kind() == reflect.String && !val.isExpr

			if val.isGt {
				operator = ">"
			} else if val.isGe {
				operator = ">="
			} else if val.isLt {
				operator = "<"
			} else if val.isLe {
				operator = "<="
			} else if val.isNe {
				operator = "!="
			}
		}

		if isMark {
			sql += fmt.Sprintf("%s %s '%v' %s ", k, operator, v, chain)
		} else {
			sql += fmt.Sprintf("%s %s %v %s ", k, operator, v, chain)
		}
	}

	return strings.TrimRight(sql, fmt.Sprintf(" %s ", chain))
}

func toInSql(list reflect.Value) string {
	sql := ""

	for i := 0; i < list.Len(); i++ {
		v := list.Index(i).Interface() // 列表中的每个值
		_, isMark := v.(string)        // 是否给值[v]加上单引号

		if val, ok := v.(Value); ok {
			vv := reflect.ValueOf(val.value)
			if val.isNot || val.isNonZero && vv.IsZero() { // 排除该元素
				continue
			}

			v = val.value // 改变 v 的值以追加到 sql
			isMark = vv.Type().Kind() == reflect.String && !val.isExpr
		}

		if isMark {
			sql += fmt.Sprintf("'%v', ", v)
		} else {
			sql += fmt.Sprintf("%v, ", v)
		}
	}

	return fmt.Sprintf("IN (%s)", strings.TrimRight(sql, ", "))
}
