package calcFunctions

import (
	"strconv"
	"strings"
)

func StringCalculator(arg ...string) []float64 {
	var resultArr = make([]float64, len(arg), len(arg))

	for i := 0; i < len(arg); i++ {

		switch {
		case strings.Contains(arg[i], "*"):

			{
				arrayStr := strings.Split(arg[i], "*")
				var floatedArr []float64
				for i := 0; i < len(arrayStr); i++ {
					newNumbers, _ := strconv.ParseFloat(arrayStr[i], 64)
					floatedArr = append(floatedArr, newNumbers)
				}

				resultArr[i] = Multiply(floatedArr)
				break
			}

		case strings.Contains(arg[i], "+"):

			{
				arrayStr := strings.Split(arg[i], "+")
				var floatedArr []float64
				for i := 0; i < len(arrayStr); i++ {
					newNumbers, _ := strconv.ParseFloat(arrayStr[i], 64)
					floatedArr = append(floatedArr, newNumbers)
				}

				resultArr[i] = Addition(floatedArr)
				break
			}
		case strings.Contains(arg[i], "/"):

			{
				arrayStr := strings.Split(arg[i], "/")
				var floatedArr []float64
				for i := 0; i < len(arrayStr); i++ {
					newNumbers, _ := strconv.ParseFloat(arrayStr[i], 64)
					floatedArr = append(floatedArr, newNumbers)
				}

				resultArr[i] = Division(floatedArr)
				break
			}

		case strings.Contains(arg[i], "-"):

			{
				arrayStr := strings.Split(arg[i], "-")
				var floatedArr []float64
				for i := 0; i < len(arrayStr); i++ {
					newNumbers, _ := strconv.ParseFloat(arrayStr[i], 64)
					floatedArr = append(floatedArr, newNumbers)
				}

				resultArr[i] = Subtraction(floatedArr)
				break
			}

		}
	}
	return resultArr
}
