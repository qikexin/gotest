package main

import (
	"fmt"
	"flag"
	"os"
	lib1 "test/example4/test2/lib"
	"test/example4/test1/lib"
)

var name string
func init(){
	flag.StringVar(&name,"name","everyone","input a username")
}
//var name = flag.String("name","everyone","input a username")
func main()  {
	flag.Usage = func(){
		fmt.Fprintf(os.Stderr,"usage of %s:\n","question")
		flag.PrintDefaults()
	}
	flag.Parse()
	lib1.Hello(name)
	lib.Say(name)
	//fmt.Printf("hello,%s!\n",name)
}