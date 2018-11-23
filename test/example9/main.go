package main

import (
	"net/http"
	"fmt"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"regexp"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func main()  {
	/*resp, err := http.Get("http://www.baidu.com")
if err !=nil {
	panic(err)
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
fmt.Println(body)*/
	/*client := &http.Client{

	}*/
	//resp,err := client.Get("http://www.baidu.com")
	/*req, err := http.NewRequest("GET","http://www.baidu.com",nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)*/
	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("response err", resp.StatusCode)
		return
	}
	e := determineEncoding(resp.Body)
	/*utf8Reader := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewReader)
	all,err :=ioutil.ReadAll(utf8Reader)*/
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	all,err :=ioutil.ReadAll(utf8Reader)
	//all,err :=ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}
	//fmt.Printf("%s 222222 \n",e)
	printCityList(all)
}

//determineEncoding函数用于解析resp.body的编码
func determineEncoding(r io.Reader) encoding.Encoding{
	bytes,err := bufio.NewReader(r).Peek(1024)
	if err  != nil{
		panic(err)
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return  e
}
func printCityList(contents []byte){
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents,-1)

	for _,m := range matches{
		fmt.Printf("City: %s, URL: %s\n",m[2],m[1])
	}
	fmt.Printf("matches found: %d\n",len(matches))
}