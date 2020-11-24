package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Print("logic 1")


	
}

func sortaSum(a int, b int) int {

	if (a + b) >= 10 && (a+b) <= 19 {

		return 20
	}

	return (a+b)
}

func alarmClock(day int, vacation  bool) string {

	if day >= 1 && day <= 5 && vacation || !vacation && day == 6 || !vacation && day == 0 {

		return "10:00"
	}

	if day >= 1 && day <= 5 && !vacation {

		return "7:00"
	}

	if day >= 1 && day <= 5 && !vacation {

		return "off"
	}

	return ""
}

func love6(a int, b int) bool {




	if (a+b) == 6 || math.Abs(float64(a) - float64(b)) == 6 {

		return true
	}

	return false
}


func int1To10(n int, outsideMode bool) bool {

	var rult bool = n >= 1 && n <= 10 && !outsideMode || outsideMode && n <= 1 || outsideMode && n >= 10

	return rult



}

func speicaleleven(n int) bool {


	var rult bool = n % 11 == 0 || (n - 1) % 11 == 0 || n + 1 % 11 == 0

	return rult
}

func more20(n int) bool {


	var rult bool = n % 3 == 0 && n % 5 != 0 || n % 3 != 0 && n % 5 == 0

	return rult


}




func old35(n int) bool {

	return n % 3 == 0 && n % 5 != 0 || n % 3 != 0 && n % 5 == 0
}

func less20(n int) bool {

	return (n+2) % 20 == 0 || (n+1) % 20 == 0
	
}

func nearTen(num int) bool {

	// all have to be on the same line

	return (num - 2) % 10 == 0 || (num - 1) % 10 == 0 || (num % 10) == 0 || (num + 2) % 10 == 0 || (num + 1) % 10 == 0
}

func teenSum(a int, b int) int {

	if a >= 13 && a <= 19 || b >= 13 && b <= 19 {

		return 19
	}

	return a + b
}



