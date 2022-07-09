package calcFunctions

import (
	"strconv"
	"strings"
)

func StringCalculator(arg ...string) []float64 {
	var resultArr = make([]float64, len(arg))

	for i := 0; i < len(arg); i++ {

		switch {
		case strings.Contains(arg[i], "*"):

			{
				arrayStr := strings.Split(arg[i], "*")
				floatedArr := changeStringToSlice(arrayStr)
				resultArr[i] = Multiply(floatedArr)
				break
			}

		case strings.Contains(arg[i], "+"):

			{
				arrayStr := strings.Split(arg[i], "+")
				floatedArr := changeStringToSlice(arrayStr)
				resultArr[i] = Addition(floatedArr)
				break
			}
		case strings.Contains(arg[i], "/"):

			arrayStr := strings.Split(arg[i], "/")
			floatedArr := changeStringToSlice(arrayStr)
			resultArr[i] = Division(floatedArr)
			fallthrough

		case strings.Contains(arg[i], "-"):

			{
				arrayStr := strings.Split(arg[i], "-")
				floatedArr := changeStringToSlice(arrayStr)
				resultArr[i] = Subtraction(floatedArr)
				break
			}

		}
	}
	return resultArr
}

//function converts slices of string into slices of floats
func changeStringToSlice(s []string) []float64 {
	var floatedArr []float64
	for i := 0; i < len(s); i++ {
		newNumbers, _ := strconv.ParseFloat(s[i], 64)
		floatedArr = append(floatedArr, newNumbers)
	}
	return floatedArr
}
