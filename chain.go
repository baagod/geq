package eq

import (
    "fmt"
    "reflect"
    "strings"
)

type Builder interface {
    SQL() string
}

// WhereSQL 返回带上 WHERE 前缀的 sql 字符串
func (m And) WhereSQL() string {
    if sql := m.SQL(); sql != "" {
        return "WHERE " + sql
    }
    return ""
}

// ---- 等于 AND ----

type And map[string]any

func (m And) SQL() string { return toSQL(m, "AND") }

func (m And) Cut() And { return cut(m) }

// ---- 或者 OR ----

type Or map[string]any

func (m Or) SQL() string { return toSQL(m, "OR") }

func (m Or) Cut() Or { return cut(m) }

// ---- 私有方法 ----

func toSQL(eq map[string]any, chain string) string {
    var sql string
    operator := "="

    for k, v := range eq {
        mvv := reflect.ValueOf(v)
        if mvv.Kind() == reflect.Slice || mvv.Kind() == reflect.Array {
            listSql := inSQL(mvv)
            sql += fmt.Sprintf("%s %s %s ", k, listSql, chain)
            continue
        }

        if val, ok := v.(*Value); ok {
            vv := reflect.ValueOf(val.value)
            if val.isOut || val.isNoZero && vv.IsZero() { // 排除该字段
                continue
            }

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
            } else if val.isBetween {
                operator = "BETWEEN"
            } else {
                operator = "="
            }
        }

        if _, ok := v.(string); ok {
            sql += fmt.Sprintf("%s %s '%v' %s ", k, operator, v, chain)
        } else {
            sql += fmt.Sprintf("%s %s %v %s ", k, operator, v, chain)
        }
    }

    return strings.TrimRight(sql, fmt.Sprintf(" %s ", chain))
}

func inSQL(list reflect.Value) string {
    sql := ""

    for i := 0; i < list.Len(); i++ {
        v := list.Index(i).Interface() // 列表中的每个值

        if val, ok := v.(*Value); ok {
            vv := reflect.ValueOf(val.value)
            if val.isOut || val.isNoZero && vv.IsZero() { // 排除该元素
                continue
            }
        }

        if _, ok := v.(string); ok {
            sql += fmt.Sprintf("'%v', ", v)
        } else {
            sql += fmt.Sprintf("%v, ", v)
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
                if vv, ok := v.(*Value); ok {
                    if vv.isOut || vv.isNoZero && reflect.ValueOf(vv.value).IsZero() { // 排除该元素
                        mv = reflect.AppendSlice(mv.Slice(0, i), mv.Slice(i+1, mv.Len())) // 移除第 i 个元素
                    }
                }
            }

            if mv.Len() == 0 {
                delete(m, k)
                continue
            }

            m[k] = mv.Interface()
        }

        if val, ok := v.(*Value); ok {
            if val.isOut || val.isNoZero && reflect.ValueOf(val.value).IsZero() { // 排除该字段
                delete(m, k)
            }
        }
    }

    return m
}
