package parser

import (
	"testing"
	"test/crawller/fatcher"
	"fmt"
)

func TestParseCityList(t *testing.T) {
	contents, err := fatcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}
	//ParseCityList(contents)
	fmt.Printf("%s\n",contents)
}