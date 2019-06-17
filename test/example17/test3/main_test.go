/*
Author: lipengwei
Date: 2019/5/6
Description: 
*/
package main

import (
	"testing"
	"fmt"
)

func TestHello(t *testing.T)  {
	var name string
	greeting,err := Hello(name)
	if err == nil {
		t.Errorf("the error is nil,but it should" +
		"not be.(name=%q)",name)
	}
	if greeting != "" {
		t.Errorf("nonempty greeting,but it should not be" +
			"(name=%q)",name)
	}
	name = "Robert"
	greeting,err = Hello(name)
	if err != nil{
		t.Errorf("the error is not nil,but it should be. (name=%q)",name)
	}
	if greeting == ""{
		t.Errorf("empty greeting,but it should not be (name=%q)",name)
	}
	expectd := fmt.Sprintf("hello,%s!",name)
	if greeting != expectd {
		t.Errorf("the actual greeting %q is not be expected (name=%q)",greeting,name)
	}
	t.Logf("the expected greeting is %q\n",expectd)
}

func TestIntroduce(t *testing.T)  {
	intro := Introduce()
	expected := "welcome to my golang column"
	if intro != expected {
		t.Errorf("the actual introduce %q is not the expected",intro)
		t.Logf("the expected introduce is %q\n",expected)
	}
}