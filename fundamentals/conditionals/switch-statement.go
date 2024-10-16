package conditionals

import (
	"fmt"
	"runtime"
	"time"
)

func SwitchExample() {
	fmt.Println("> switchExample")
	// the break statement is provided automatically
	fmt.Println("Go runs on:")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}
}

func SwitchEvaluationOrder() {
	fmt.Println("> switchEvaluationOrder")
	today := time.Now().Weekday()
	fmt.Printf("today: %v %v\n", today, int(today))

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away!")
	}
}

func SwitchWithNoCondition() {
	fmt.Println("> switchWithNoCondition")
	t := time.Now().Hour()

	switch {
	case t < 12:
		fmt.Println("Morning")
	case t < 18:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Night")
	}
}
