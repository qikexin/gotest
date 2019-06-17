package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	dir := "D:\\software\\gotest\\src\\gotest\\test\\example14\\test1"
	listAll(dir,0)
}
func listAll(path string, curHier int){
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil{fmt.Println(err); return}
	/*for i,j := range fileInfos{
		fmt.Println(i, "=" ,j.Name())
	}
	fmt.Println("以上输出所有目录")*/
	for _, info := range fileInfos{
		if info.IsDir(){
			for tmpHier := curHier; tmpHier > 0; tmpHier--{
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(),"\\")
			listAll(path + "\\" + info.Name(),curHier + 1)
		}else{
			for tmpHier := curHier; tmpHier > 0; tmpHier--{
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}