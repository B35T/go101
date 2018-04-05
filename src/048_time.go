package main

import (
	"fmt"
	t "time"
)

func main() {
	p := fmt.Println

	now := t.Now()
	p(now)

	var then = t.Date(1994,01,11,23,12,1,12234343,t.UTC)
	p(then)

	p(then.Minute())
	p(then.Hour())
	p(then.Day())
	p(then.Month())
	p(then.Year())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.ISOWeek())

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Nanoseconds())
	p(diff.Seconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
