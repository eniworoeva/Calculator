package calcFunctions

import (
	"fmt"
	"testing"
)

func TestStringCalculator(t *testing.T) {
	var nums = []struct {
		inp  []string
		outp []float64
	}{
		{[]string{"500-450-65-85", "12+55+89+100", "1500/3/2/2", "3*67*2*2"},
			[]float64{-100, 256, 125, 804}},
		{[]string{"600-150-85-85", "120+55+45+100", "1200/4/3/2", "4*12*7*5"},
			[]float64{280, 320, 50, 1680}},
	}
	for _, inputs := range nums {
		floatResult := StringCalculator(inputs.inp...)
		for i := 0; i < len(inputs.outp); i++ {
			if floatResult[i] != inputs.outp[i] {
				t.Errorf("Expected output %v received %v", inputs.outp, floatResult)
				fmt.Println(floatResult[i])
			}
		}
	}
}

func TestDivision(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{5000, 5, 2, 5}, 100},
		{[]float64{200, 2, 2, 2}, 25},
	}
	for _, inputs := range nums {
		floatResult := Division(inputs.inp)
		if floatResult != inputs.outp {
			t.Errorf("Division: Expected Output %v : output %v", inputs.outp, floatResult)
		}
	}
}
func TestMultiply(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{3, 7, 15, 11}, 3465},
		{[]float64{43, 2, 1, 3}, 258},
	}
	for _, inputs := range nums {
		floatResult := Multiply(inputs.inp)
		if floatResult != inputs.outp {
			t.Errorf("Multiply: Expected Output %v: output %v", inputs.outp, floatResult)
		}
	}
}

func TestSubtraction(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{15, 5, 2, 2}, 6},
		{[]float64{20, 3, 2, 1}, 14},
	}
	for _, inputs := range nums {
		floatResult := Subtraction(inputs.inp)
		if floatResult != inputs.outp {
			t.Errorf("Subtraction: Expected Output %v : output %v", inputs.outp, floatResult)
		}
	}
}
func TestAddition(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{23, 2, 7, 16}, 48},
		{[]float64{65, 21, 16, 70}, 172},
	}
	for _, inputs := range nums {
		floatResult := Addition(inputs.inp)
		if floatResult != inputs.outp {
			t.Errorf("Addition: Expected Output %v: output %v", inputs.outp, floatResult)
		}
	}
}
