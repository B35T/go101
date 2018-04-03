package main

import (
	"fmt"
	"time"
)

func main() {
	const i = 1
	switch i {
	case 1:
		fmt.Println("one")
		break
	case 2:
		fmt.Println("two")
		break
	case 3:
		fmt.Println("three")
		break
	default:
		fmt.Println("other")
		break
	}

	switch time.Now().Weekday() {
	case time.Sunday:
		fmt.Println("sunday")
	case time.Monday:
		fmt.Println("monday")
	case time.Thursday:
		fmt.Println("thursday")
	case time.Wednesday:
		fmt.Println("wednesday")
	case time.Tuesday:
		fmt.Println("tuesday")
	case time.Friday:
		fmt.Println("friday")
	case time.Saturday:
		fmt.Println("saturday")
	}

	var t = time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 12:
		fmt.Println("It's after noon")
	}

	whatAM := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am bool")
		case string:
			fmt.Println("i am string")
		case int:
			fmt.Println("i am int")
		default:
			fmt.Println("not i am", t)
		}
	}

	whatAM("name")
	whatAM(24)
	whatAM(true)
}
