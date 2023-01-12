package ged

import "fmt"

type Cond struct {
	list      []Eq
	operators []string
}

func (c Cond) AND(m Eq) Cond {
	c.list = append(c.list, m)
	c.operators = append(c.operators, "AND")
	return c
}

func (c Cond) OR(m Eq) Cond {
	c.list = append(c.list, m)
	c.operators = append(c.operators, "OR")
	return c
}

func (c Cond) SQL() string {
	sql, length := "", len(c.list)
	for i, v := range c.list {
		sql += fmt.Sprintf("(%s)", v.SQL())
		if i+1 < length {
			sql += fmt.Sprintf("\n  %s ", c.operators[i])
		}
	}
	return sql
}
