package parser

import (
	"test/crawller/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult  {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	limit := 10
	for _,m := range matches{
		//fmt.Printf("City: %s, URL: %s\n",m[2],m[1])
		result.Items = append(result.Items,"City" + string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: engine.NilParser,
		})
		limit--
		if limit == 0{
			break
		}
	}
	return  result
	//fmt.Printf("matches found: %d\n",len(matches))
}