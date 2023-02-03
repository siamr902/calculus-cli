package main

import (
	"fmt"
	"log"
	"flag"
	"math/rand"
	"regexp"
	"time"
	clc "github.com/TheDemx27/calculus"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	equation := flag.String("e", "(x^2+1)/2", "Set the equation for differentiating or integrating.")
	mode := flag.String("m", "derivative", "Set the mode to `derivative` or `integral`.")
	diffPoint := flag.Float64("x", 1, "Differentiate the equation at a value.")
	upperBound := flag.Float64("u", 20, "Set the upper bound for the definite integral.")
	lowerBound := flag.Float64("l", 0, "Set the lower bound for the definite integral.")
	flag.Parse()

	r, err := regexp.Compile("^[A-Za-z0-9*/^+-/(/)]+$")
	if err != nil {
		log.Fatal("Invalid Expression")
	}

	// for accurate results, enclose variables within parentheses (e.g. 3x -> 3(x) or 9x^3 -> 9(x^3))
	matched := r.Match([]byte(*equation))
	
	if matched == false {
		log.Fatal("Equation contains invalid characters.")
	}

	switch *mode {
	case "derivative":
		ans := evaluateDerivative(equation, diffPoint)
		fmt.Printf("The derivative at the point %v is %v", *diffPoint, ans)
	case "integral":
		ans := evaluateIntegral(equation, upperBound, lowerBound)
		fmt.Printf("The integral from %v to %v is %v", *lowerBound, *upperBound, ans)
	}
	
}

func evaluateDerivative(equation *string, point *float64) float64 {
	f := clc.NewFunc(*equation)
	ans := f.Diff(*point)
	return ans
}

func evaluateIntegral(equation *string, upper *float64, lower *float64) float64 {
	f := clc.NewFunc(*equation)
	ans := f.AntiDiff(*lower, *upper)
	return ans
}