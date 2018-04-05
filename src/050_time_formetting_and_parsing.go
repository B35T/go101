package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()

	p(t.Format(time.RFC3339))

	t1, _:= time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)
	p(t1)

	p(t.Format("3:04AM"))
	p(t.Format("Mon Jan _2 12:12:01 2006"))

	from := "3 04 PM"
	t2, e := time.Parse(from,"8 41 AM")
	fmt.Println(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}