package repository

import (
	"time"
)

func DateConv(t time.Time) time.Time {
	var layout = "02.01.2006"
	t1 := t.Format(layout)

	tConv, _ := time.Parse(layout, t1)

	return tConv
}

func concantenateProducts(sl []string, slp []string) []string {
	for _, v := range sl {
		slp = append(slp, v)
	}
	return slp
}
