package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	
	fmt.Println("learning go ");


	fmt.Printf("Sleep In: %v \n", sleepIn(false, false))

	fmt.Printf("monkey trouble:  %v \n", monkeyTrouble(true,true))

	fmt.Printf("sumDouble: %v \n", sumDouble(1,2))

	fmt.Printf("diff21: %v \n", diff21(10))

	fmt.Printf("parrotTrouble: %v \n", parrotTrouble(true, 6))

	fmt.Printf("posNeg: %v \n", posNeg(1, -1, false))

	fmt.Printf("notString: %v \n", notString("candy"))

	fmt.Printf("missingChars: %v \n", missingChar("kitten",1))


	fmt.Printf("frontBack: %v \n",frontBack("code"))


	fmt.Printf("front3 %v \n", front3("Java"))

	fmt.Printf("backArround %v \n", backAround("Hello"))

	fmt.Printf("or35 %v \n", or35(8))


	fmt.Printf("hasTeen %v \n", hastTeen(13,20,10))

	fmt.Printf("delDel: %v \n", delDel("adelbc"))

	fmt.Printf("mixStart %v \n", mixStart("mix snacks"))

	fmt.Printf("startOz: %v \n", startOz("zzzz"))













}

func sleepIn(weekday bool, vacation bool)  bool {

	return !weekday || vacation;

	
}

func monkeyTrouble(aSmile bool, bSmile bool) bool {

	return aSmile == true && bSmile == true
}

func sumDouble(a int, b int) int  {

	if a == b {

		return (a + b) * 2

	}

	return a + b;

}

func diff21(n int) int {

	var rult float64

	if n > 21 {

		rult = math.Abs(float64(n - 21)) * 2



	} else {

		rult = math.Abs(float64(n -21))
	}

	//var fnal int = int(rult);
	//
	//return fnal;

	var fnal = int(rult)

	return fnal






}

func parrotTrouble(talking bool, hour int) bool {

	var rult bool;

	if talking && hour < 7 || talking && hour > 20 {

		rult =  true

	}

	return rult;


}

func makes10(a int, b int) bool {

	if a+b == 10 || a == 10 || b == 10 {

		return true
	}

	return false
}


func nearHundred(n int) bool  {



	if math.Abs(float64(100-n)) <= 10 || math.Abs(float64(200-n)) <= 10 {

		return true

	}

	return false
}

func posNeg(a int, b int, negative bool) bool  {

	if negative && a <= -1 && negative && b <= -1 || a <= -1 && b >= 0 || b <= -1 && a >= 0 {

		return true

	}

	return false

}

func notString(str string) string  {

	if len(str) > 2 && str[0:3] == "not" {

		return str;
	}

	return "not " + str

}

func missingChar(str string, n int) string {


	return str[0:n] + str[n+1:len(str)]
}

func frontBack(str string) string {


	rult := str[len(str)-1:len(str)] + str[1:len(str)-1] + str[0:1]

	return rult
}


func front3(str string) string {

	var rult string
	if str == "" || len(str) <= 3 {

		return str + str + str

	}

	for i := 0; i < 3; i++ {

		rult += str[0:3]
	}

	return rult


}

func backAround(str string) string  {

	if str == "" {

		return str
	}

	ch := str[len(str)-1]


	rult := string(ch) + str + string(ch)

	return rult

}

func or35(n int) bool {


	return n % 3  == 0 || n % 5 == 0
}

func front22(str string) string {

	if str == "" || len(str) <= 2 {

		return str + str + str
	}

	rult := str[0:3]

	return rult+rult+rult
}

func startHi(str string) bool {

	if str == "" || len(str) < 2 {

		return false
	}

	return str[0:2] == "hi"

}

func icyHot(temp1 int, temp2 int) bool {

	if temp1 < 0 && temp2 > 100 || temp2 < 0 && temp1 > 100 {

		return true

	}

	return false
}

func in1020(a int, b int) bool {

	if a >= 10 && a <= 20 || b >= 10 && b <= 20 {

		return true
	}

	return false
}

func hastTeen(a int, b int, c int) bool {

	if a >= 13 && a <= 19 {

		return true

	} else if b >= 13 && b <= 19 {

		return true
	} else if c >= 13 && c <= 19 {

		return true
	}

	return false


}

func delDel(str string) string {

	if str == "" || len(str) < 3 || strings.Index(str,"del") == -1 {

		return str
	}

	var rult string

	if strings.Index(str, "del") == 1 {

		rult = str[0:1] + str[len("del")+1:len(str)]
	}

	return rult
}

func mixStart(str string) bool {

	if str == "" || len(str) < 3 {

		return false
	}

	if strings.Index(str, "ix") == 1 {

		return true


	}

	return false
}

func startOz(str string) string {

	var rult string

	if str == ""  {

		return ""

	}

	if strings.Index(str, "o") == 0 && strings.Index(str, "z") == 1 {

		rult = "oz"
	}

	if strings.Index(str, "o") == 0 && strings.Index(str, "z") == -1 {

		rult = "o"

	}

	if str[0:1] != "o" && str[1:2] == "z" {

		rult = "z"
	}

	return rult
}




