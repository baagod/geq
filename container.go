package eq

import (
    "fmt"
)

type Container struct {
    list      []Builder
    operators []string
}

func (c Container) AND(m Builder) Container {
    c.list = append(c.list, m)
    c.operators = append(c.operators, "AND")
    return c
}

func (c Container) OR(m Builder) Container {
    c.list = append(c.list, m)
    c.operators = append(c.operators, "OR")
    return c
}

func (c Container) ToSQL() string {
    var sql string
    length := len(c.list)
    for i, v := range c.list {
        sql += fmt.Sprintf("(%s)", v.SQL())
        if i+1 < length {
            sql += fmt.Sprintf("\n  %s ", c.operators[i+1])
        }
    }
    return sql
}
