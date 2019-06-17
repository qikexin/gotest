/*
Author: lipengwei
Date: 2019/5/6
Description: 
*/
package test4

import (
	"math"
)

func GetPrimes(max int) []int  {
	if max < 1 {
		return []int{}
	}
	marks := make([]bool,max)
	var count int
	squareRoot := int(math.Sqrt(float64(max)))
	for i := 2;i <= squareRoot;i++ {
		if marks[i] == false {
			for j := i*i;j<max;j+=1 {
				if marks[i] == false{
					marks[j] = true
					count++
				}
			}
		}
	}
/*	fmt.Println(max-count)
	fmt.Println("\n")*/
	primes := make([]int,0,count-max)
	for i := 2;i < max; i++ {
		if marks[i] == false {
			primes = append(primes,i)
		}
	}
	return primes
}