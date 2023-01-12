package ged

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"testing"
)

func TestName(t *testing.T) {
	m := Eq{
		"create_at": Between(10, 12),
		"name":      V("HJZ"),
		"age":       Expr("55"),
		"num":       []any{1, Expr("55"), 3},
	}

	// a = 1 || (b = 2 && c = 3)
	// (a = 1 && b = 2) || c = 3 || c = 4
	// (a = 1 && b = 2) || c = 3 && d = 4

	fmt.Println(m.SQL())

	c := Eq{"a": 1}.OR(Eq{"b": 2, "c": 3})
	c = Eq{"a": 1, "b": 2}.OR(Eq{"c": 3, "d": 4})
	fmt.Println("c.SQL:", c.SQL())

	//e := sq.Eq{"col1": 1, "col2": 2}
	//fmt.Println(e.ToSql())
	//
	s := sq.Or{
		sq.Eq{"col1": 1, "col2": 2},
		sq.Eq{"col1": 3, "col2": 4},
	}

	//fmt.Println(s.ToSql())
}
