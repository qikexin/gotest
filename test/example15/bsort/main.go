/*
Author: lipengwei
Date: 2018/12/4
Description: 
*/
package main

import "fmt"

func bsort(a []int)  {
	for i := 0;i < len(a)-1;i++{
		for j := 0;j<len(a)-i-1;j++{
			if a[j] < a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
}
func main()  {
	b := [...]int{8,5,4,6,2,9,1,3}
	bsort(b[:])
	fmt.Println(b)
}