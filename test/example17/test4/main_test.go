/*
Author: lipengwei
Date: 2019/5/6
Description: 
*/
package test4

import "testing"

func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N;i++ {
		GetPrimes(1000)
	}
}