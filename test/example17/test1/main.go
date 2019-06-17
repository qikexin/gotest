/*
Author: lipengwei
Date: 2019/4/30
Description: 
*/
package main

import "fmt"

/*
func main()  {
	var count uint32
	trigger := func (i uint32,fn func()){
		for {
			if n := atomic.LoadUint32(&count);n == i{
				fn()
				atomic.AddUint32(&count,1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}

	for i := uint32(0);i<10;i++{
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i,fn)
		}(i)
	}
	trigger(10, func(){})
}*/
func main()  {
	//number1 := []int{1,2,3,4,5,6}
	/*number2 := [...]int{1,2,3,4,5,6}
	maxIndex2 := len(number2) - 1
	fmt.Println(maxIndex2)
	number3 := number2
	for i,e := range number2 {
		if i == maxIndex2 {
			number2[0] += e
		}else {
			number2[i+1] += e
		}
	}
	fmt.Println(number2)
	fmt.Println(number3)*/
	/*for i := range number1 {
		fmt.Println(number1[i])
		if  i == 3 {
			number1[i] |= i
		}
	}*/
	//fmt.Println(number1)

	/*value1 := [...]int{0, 1, 2, 3, 4, 5, 6}
	switch 1 + 3 {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[2], value1[3]:
		fmt.Println("2 or 3")
	case value1[4], value1[5], value1[6]:
		fmt.Println("4 or 5 or 6")
	}*/
	/*value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[4] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}*/

	/*value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value3[4] {
	case 0, 1, 1+1:
		fmt.Println("0 or 1 or 2")
	case 2, 3, 4:
		fmt.Println("2 or 3 or 4")
	case 5, 6:
		fmt.Println("4 or 5 or 6")
	}*/

	value6 := interface{}(byte(127))
	switch t:= value6.(type) {
	case uint32,uint16:
		fmt.Println("uint8/uint16")
	case byte:
		fmt.Println("byte")
	default:
		fmt.Printf("unsupport type: %T",t)
	}
}