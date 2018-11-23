package engine

import (
	"test/crawller/fatcher"
	"log"
)
func Run(seeds ...Request){  //seeds表示是一个request类型的可变参数
	var requests []Request
	for _, r := range seeds{
		requests = append(requests,r)
	}

	for len(requests)> 0{
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching %s",r.Url)
		body, err := fatcher.Fetch(r.Url)
		if err != nil{
			log.Printf("Fetcher: error " + "fetching url %s : %v",r.Url,err)
			continue
		}

		paseResult := r.ParserFunc(body)
		requests = append(requests,paseResult.Requests...)
		for _,item := range paseResult.Items{
			log.Printf("got item %s", item)
		}
	}
}
