//
// EPITECH PROJECT, 2020
// 201yams_2019
// File description:
// error
//

package functions

import (
	"fmt"
	"math"
)

//BetweenIQ calculate of people have an IQ between min IQ and max IQ
func BetweenIQ(numbers []int) int {
	res := 0.0

	for i := int(1); i < 4; i++ {
		if numbers[i] < 1 || numbers[i] > 200 {
			return 84
		}
	}
	if numbers[2] > numbers[3] {
		return 84
	}
	for x := float64(numbers[2]) * 100; x != float64(numbers[3])*100; x++ {
		res += (1. / (float64(numbers[1]) * math.Sqrt(2.*math.Pi))) *
			math.Exp(-math.Pow((x/100)-float64(numbers[0]), 2)/(2.*math.Pow(float64(numbers[1]), 2)))
	}
	fmt.Printf("%.1f%% of people have an IQ between %d and %d\n", res, numbers[2], numbers[3])
	return 0
}

//InfIQ Calculate of people have an IQ inferior of min IQ
func InfIQ(numbers []int) int {
	res := 0.0
	
	for i := int(1); i < 3; i++ { 
		if numbers[i] < 1 || numbers[i] > 200 {
			return 84
		}
	}
	for x := 0.0; x != float64(numbers[2])*100; x++ {
		res += (1. / (float64(numbers[1]) * math.Sqrt(2.*math.Pi))) *
			math.Exp(-math.Pow((x/100)-float64(numbers[0]), 2)/(2.*math.Pow(float64(numbers[1]), 2)))
	}
	fmt.Printf("%.1f%% of people have an IQ inferior to %d\n", res, numbers[2])
	return 0
}

//DataDeviation  p(x) = 1/(σ √(2 π)) exp( - (x - μ)² / σ²)
func DataDeviation(numbers []int) int {
	if numbers[1] < 1 || numbers[1] > 200 {
		return 84
	}
	for x := 0.0; x != 201.0; x++ {
		res := (1. / (float64(numbers[1]) * math.Sqrt(2.*math.Pi))) *
			math.Exp(-math.Pow(x-float64(numbers[0]), 2)/(2.*math.Pow(float64(numbers[1]), 2)))
		fmt.Printf("%d %.5f\n", int(x), res)
	}
	return 0
}

//MathParse args
func MathParse(numbers []int) int {
	switch len(numbers) {
	case 2:
		return DataDeviation(numbers)
	case 3:
		return InfIQ(numbers)
	case 4:
		return BetweenIQ(numbers)
	default:
		return 84
	}
}
