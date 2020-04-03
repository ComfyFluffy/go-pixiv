package pixiv

import (
	"fmt"
	"testing"
)

func assert(e bool, x ...interface{}) {
	if !e {
		panic(fmt.Sprint("assert failed ", x))
	}
}

func TestDate(t *testing.T) {
	d := Date("2020-04-03")
	dn := NewDate(2020, 4, 3)
	assert(d == dn, d, dn)
	t.Log(dn)

	y := d.Year()
	assert(y == 2020, y, 2020)

	m := d.Month()
	assert(m == 4, m, 4)

	dd := d.Day()
	assert(dd == 3, d, 3)
}
