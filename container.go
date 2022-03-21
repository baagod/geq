package eq

import (
	"fmt"
)

type Container struct {
	list      []Builder
	operators []string
}

func (c Container) AND(eq Builder) Container {
	c.list = append(c.list, eq)
	c.operators = append(c.operators, "AND")
	return c
}

func (c Container) OR(eq Builder) Container {
	c.list = append(c.list, eq)
	c.operators = append(c.operators, "OR")
	return c
}

func (c Container) ToSQL() string {
	var sql string
	length := len(c.list)
	for i, v := range c.list {
		sql += fmt.Sprintf("(%s)", v.ToSql())
		if i+1 < length {
			sql += fmt.Sprintf("\n  %s ", c.operators[i+1])
		}
	}
	return sql
}

func (c Container) Cut() Container {
	for _, v := range c.list {
		v.Cut()
	}
	return c
}

func (c Container) String() string {
	return fmt.Sprint(c.list)
}
