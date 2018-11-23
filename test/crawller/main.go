package main

import (
	"test/crawller/engine"
	"test/crawller/zhenai/parser"
)

func main()  {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

