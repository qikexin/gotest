package parser

import (
	"test/crawller/engine"
	"regexp"
)

//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9a-zA-Z)"`

const cityRe = `<a href="http://album.zhenai.com/u/([0-9]+)"[^>]*>([^<]+)</a>`
func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches{
		name := string(m[2])
		result.Items = append(result.Items,"User "+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return  ParseProfile(c,name)
			},
			//ParserFunc: ParseProfile,
		})
		limit--
		if limit == 0{
			break
		}
	}

	return result
}