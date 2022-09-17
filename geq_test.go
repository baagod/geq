package eq

import "testing"

func TestAnd(t *testing.T) {
    and := And{
        "lid":   11,
        "state": 1,
        "msg":   "test msg",
        "date":  Between("2022-01", "2022-02"),
    }

    println(and.SQL())
}
