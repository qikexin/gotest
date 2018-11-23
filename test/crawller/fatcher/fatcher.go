package fatcher

import (
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"regexp"
	"fmt"
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Fetch(url string) ([]byte,error){
	resp, err := http.Get(url)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("response err", resp.StatusCode)
		return nil,fmt.Errorf("wrong status code, %d",resp.StatusCode)
	}
	e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

//determineEncoding函数用于解析resp.body的编码
func determineEncoding(r io.Reader) encoding.Encoding{
	bytes,err := bufio.NewReader(r).Peek(1024)
	if err  != nil{
		log.Printf("fetcher error: %v",err)
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return  e
}
func printCityList(contents []byte){
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`)
	//matchs := re.FindAll(contents,-1)
	matches := re.FindAllSubmatch(contents,-1)
	fmt.Println(matches)
	/*for _,m := range matches{
		fmt.Printf("City: %s, URL: %s\n",m[2],m[1])
	}*/
	fmt.Printf("matches found: %d\n",len(matches))
}

