/*
Author: lipengwei
Date: 2019/4/28
Description: 
*/
package main

import (
	"fmt"
	"github.com/pkg/errors"
)


type Printer func(contents string) (n int,err error)
/*
只要两个函数的参数列表和结果列表中元素顺序及其类型是一致的，我们就可以说他们是一样的函数，或者说他们是实现了同一个函数类型的函数
*/
func printTostd(contents string) (bytesNum int,err error){
	return fmt.Println(contents)
}
type operate func(x,y int) int

func calcculate(x,y int,op operate)(int,error) {
	if op == nil {
		return 0,errors.New("invalid operation")
	}
	return op(x,y),nil
}

/*func genCalculator(op operate) calculateFunc {
	return func(x,y int) (int,err) {
		if op == nil {
			return 0,errors.New("invalid operation")
		}
		return op(x,y),nil
	}
}*/
/*
所有传递给函数的参数都会被复制，函数内部使用的并不是参数的原值，而是它的副本。由于数组是值类型，所以每一次复制都会拷贝它以及他的所有元素值，函数在
内部修改的是原数组的拷贝副本，原数组并没有改变
*/
func modifyArray(a [3]string) [3]string{
	a[1] = "x"
	return a
}
func main()  {
	op := func(x,y int) int {
		return x+y
	}
	result,err :=calcculate(30,20,op)
	if err == nil {
		fmt.Printf("result=%d",result)
		fmt.Println("result= ",result)
	}

	array1 := [3]string{"a","b","c"}
	fmt.Printf("the array: %v\n",array1)
	array2 := modifyArray(array1)
	fmt.Printf("the modified array:%v\n",array2)
	fmt.Printf("the original array:%v\n",array1)


}