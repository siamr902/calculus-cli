package derivative

import (
	clc "github.com/TheDemx27/calculus"
)

func EvaluateDerivative(equation *string, point *float64) float64 {
	f := clc.NewFunc(*equation)
	ans := f.Diff(*point)
	return ans
}
