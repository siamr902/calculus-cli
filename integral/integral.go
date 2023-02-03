package integral

import (
	clc "github.com/TheDemx27/calculus"
)

func EvaluateIntegral(equation *string, upper *float64, lower *float64, averageValue *bool) float64 {
	f := clc.NewFunc(*equation)
	ans := f.AntiDiff(*lower, *upper)

	if *averageValue {
		return ans / (*upper - *lower)
	}
	return ans
}
