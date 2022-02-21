package calcFunctions

import (
	"fmt"
	"strconv"
)

func Multiply(arr []float64) float64 {

	var multiplyContainer float64 = 1
	for _, val := range arr {
		multiplyContainer = multiplyContainer * val
	}
	multiplyContainer, _ = strconv.ParseFloat(fmt.Sprintf("%0.6f", multiplyContainer), 64)
	return multiplyContainer
}

func Addition(arr []float64) float64 {
	var additionContainer float64 = 0
	for _, val := range arr {
		additionContainer = additionContainer + val
	}
	return additionContainer
}

func Division(arr []float64) float64 {
	var divisionContainer float64 = arr[0]
	for i := 1; i < len(arr); i++ {
		divisionContainer = divisionContainer / arr[i]
	}
	divisionContainer, _ = strconv.ParseFloat(fmt.Sprintf("%f", divisionContainer), 64)
	return divisionContainer
}

func Subtraction(arr []float64) float64 {
	var subtractionContainer = arr[0]
	for i := 1; i < len(arr); i++ {
		subtractionContainer = subtractionContainer - arr[i]
	}
	return subtractionContainer
}
