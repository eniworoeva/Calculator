package main

import (
	"fmt"
	"week-2-eniworoeva/calcFunctions"
)

func main() {
	arr := calcFunctions.StringCalculator("3*5*7*4", "56+100+34+78")
	fmt.Println(arr)
	fmt.Println(calcFunctions.Multiply([]float64{20, 11, 5, 21}))
	fmt.Println(calcFunctions.Division([]float64{11, 1, 23, 6}))
	fmt.Println(calcFunctions.Addition([]float64{20, 11, 5, 21}))
	fmt.Println(calcFunctions.Subtraction([]float64{20, 11, 5, 3}))
}
