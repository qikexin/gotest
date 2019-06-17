/*
Author: lipengwei
Date: 2019/6/12
Description: 
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/astaxie/beego/logs"
)

func check(e error)  {
	if e != nil{
		panic(e)
	}
}
var (
	h bool
	v, V bool
	t, T bool
	q  *bool
	s string
	p string
	c string
	g string
)

func init()  {
	flag.BoolVar(&h,"h",false,"this help")

	flag.BoolVar(&v,"v",false,"show version and exit")
	flag.BoolVar(&V,"V",false,"show version and configure optons then exit")

	flag.BoolVar(&t,"t",false,"test configuration and exit")
	flag.BoolVar(&T,"T",false,"test configuration,dump it and exit")

	flag.Usage = usage
}
func usage()  {
	fmt.Fprintf(os.Stdout,`nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]
Options:
`)
	flag.PrintDefaults()
}
func main()  {
	logfile := "/root/test.log"
	loger := logs.NewLogger()
	loger.SetLogger(logs.AdapterFile,`{	
		"filename": `+logfile+`,
		"level": 7,
		"maxlines": 0,
		"maxsize": 0,
		"delay": true,
		"maxdays": 10,
		"rotate": true,
		"perm": 0644,
}`)
	/*if h {
		flag.Usage()
	}*/
	/*b,err := ioutil.ReadFile("a.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	fmt.Print(b)
	str := string(b)
	fmt.Println(str)*/

	/*d1 := []byte("hello golang")
	err := ioutil.WriteFile("b.txt",d1,0644)
	if err != nil {
		panic(err)
	}*/

	/*fi, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	pwd,_ := os.Getwd()
	fmt.Println(pwd)*/

	/*f, err := os.Open("/tmp/dat")
	check(err)

	b1 := make([]byte,5)
	n, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n",n,string(b1))

	o2, err := f.Seek(6,0)
	check(err)
	b2 := make([]byte,2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n",n2,o2,string(b2))

	o3,err := f.Seek(6,0)
	check(err)
	b3 := make([]byte,2)
	n3, err := io.ReadAtLeast(f,b3,2)
	check(err)
	fmt.Printf("%d bytes @%d :%s\n",n3,o3,string(b3))

	r4 := bufio.NewReader(f)
	b4,err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n",string(b4))
	f.Close()*/

	/*f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	d2 := []byte{112,334,1124,52}
	n2,err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n",n2)

	n3,err := f.WriteString("write\n")
	fmt.Printf("wrote %d bytes\n",n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4,err := w.WriteString("buffered\n")
	fmt.Printf("worte %d bytes\n",n4)
	w.Flush()*/


}