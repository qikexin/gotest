/*
Author: lipengwei
Date: 2019/5/5
Description: 
*/
package main

import (
	"github.com/pkg/errors"
	"fmt"
	"os"
	"os/exec"
)

func echo(request string)(response string,err error)  {
	if request == "" {
		err = errors.New("empty request")
	}
	response = fmt.Sprintf("echo %s",request)
	return
}

func underlyingError(err error) error  {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err

	}
	return err
}

func caller1()  {
	fmt.Println("enter function caller1.")
	caller2()
	fmt.Println("exit function caller1")
}

func caller2()  {
	fmt.Println("enter function caller2")
	s1 := []int{0,1,2,3,4}
	e5 := s1[4]
	panic("aa")
	_ = e5
	fmt.Println("exit function caller2")
}
func main()  {
	/*for _,req := range []string{"","hello!"} {
		fmt.Printf("request: %s\n",req)
		resp,err := echo(req)
		if err != nil {
			fmt.Printf("errors: %s\n",err)
			continue
		}
		fmt.Printf("response: %s\n",resp)
	}*/
	
	/*fmt.Println("enter function main")
	caller1()
	fmt.Println("exit function main")*/

	fmt.Println("enter function main")
	defer func() {
		fmt.Println("enter defer function")
		if p:=recover();p != nil {
			fmt.Printf("panic: %s\n",p)
		}
		fmt.Println("exit defer function")
	}()
	panic(errors.New("something wrong"))
	fmt.Println("exit function main")

	/*defer fmt.Println("first defer")
	for i := 0;i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n",i)
	}
	defer fmt.Println("last defer")*/
}