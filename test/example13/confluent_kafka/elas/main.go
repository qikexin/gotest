package main

import (
	"github.com/olivere/elastic"
	"log"
	"os"
	"context"
	"fmt"
	"reflect"
	"encoding/json"
)

var client *elastic.Client
var host = "http://10.9.1.213:9200/"
var ctx = context.Background()
type Employee struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age    	  int    `json:"age"`
	About     string  `json:"about"`
	Interests []string `json:"interests"`
}

func checkErr(err error)  {
	if err != nil{
		panic(err)
	}
}

//初始化
func init()  {
	errorlog := log.New(os.Stdout,"app ",log.LstdFlags)
	var err error
	client,err  = elastic.NewClient(elastic.SetErrorLog(errorlog),elastic.SetURL(host))
	if err != nil{
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("elasticsearch returned with code %d and version %s\n",code,info.Version.Number)
	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("elasticsearch version %s\n",esversion)
}
func create()  {
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	put1, err := client.Index().
		Index("lpwtest").
		Type("employee").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nindexed tweet %s to index%s,type %s\n",put1.Id,put1.Index,put1.Type)

	e2  :=  `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().
		Index("lpwtest").Type("employee").Id("2").BodyJson(e2).Do(context.Background())
	if err != nil{
		panic(err)
	}
	fmt.Printf("\nadd %s to index: %s ,type: %s\n",put2.Id,put2.Index,put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	//删除不存在的文档会报异常
	put3, err := client.Index().
		Index("lpwtest").
		Type("employee").
		Id("3").
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)
	}
func delete()  {
	res, err := client.Delete().Index("lpwtest").Type("employee").Id("2").Do(context.Background())
	checkErr(err)
	fmt.Printf("delete document id %s from index: %s type: %s\n",res.Id,res.Index,res.Type)
}
func update()  {
	res1, err := client.Update().Index("lpwtest").Type("employee").Id("1").Doc(map[string]interface{}{"age":89}).Do(context.Background())
	checkErr(err)
	fmt.Printf("update %s in index: %s type:%s  is succ ",res1.Id,res1.Index,res1.Type)
	//更新不存在的文档会报异常
	res2, err := client.Update().Index("lpwtest").Type("employee").Id("2").Doc(map[string]interface{}{"age":89}).Do(context.Background())
	checkErr(err)
	fmt.Printf("update %s in index: %s type:%s  is succ ",res2.Id,res2.Index,res2.Type)
	}
func get()  {
	get2 ,err := client.Get().Index("lpwtest").Type("employee").Id("1").Do(context.Background())
	checkErr(err)
	fmt.Printf("index: %s type:%s document %s is exist\n",get2.Index,get2.Type,get2.Id)
	//查询不存在文档会抛出异常
	get1 ,err := client.Get().Index("lpwtest").Type("employee").Id("2").Do(context.Background())
	checkErr(err)
	fmt.Printf("index: %s type:%s document %s is exist\n",get1.Index,get1.Type,get1.Id)

}
func printEmployee(res *elastic.SearchResult,err error)  {
	if err != nil{
		print(err.Error())
		return
	}
	//checkErr(err)
	var typ Employee
	for _,item := range res.Each(reflect.TypeOf(typ)){
		t := item.(Employee)
		fmt.Printf("%+v\n",t)
	}
}
func query()  {
	var res  *elastic.SearchResult
	var err error
	//取所有命中的文档
	res, err = client.Search("lpwtest").Type("employee").Do(ctx)
	printEmployee(res,err)
	fmt.Println("取所有文档完成")
	//字段相等
	q := elastic.NewQueryStringQuery("last_name:Smith")
	res,err = client.Search("lpwtest").Type("employee").Query(q).Do(ctx)
	printEmployee(res, err)
	fmt.Println("根据字段相等获取取文档完成")

	if res.Hits.TotalHits > 0{
		fmt.Printf("found a total of %d employee \n",res.Hits.TotalHits)
		for _,hit := range res.Hits.Hits {
			var t Employee
			err := json.Unmarshal(*hit.Source,&t)
			if err != nil {
				fmt.Println("deserialization failed")
			}
			fmt.Printf("employe name %s : %s \n",t.FirstName, t.LastName)
		}
	}else {
		fmt.Println("found no employee")
	}

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name","smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("lpwtest").Type("employee").Query(q).Do(ctx)
	fmt.Println("年龄大于30岁的: 输出")
	printEmployee(res,err)

	//短语搜索 搜索about字段中有rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about","rock climbing")
	res, err = client.Search("lpwtest").Type("employee").Query(matchPhraseQuery).Do(ctx)
	fmt.Println("搜索about字段中有rock climbing: 输出")
	printEmployee(res,err)

	//分析interests
	/*aggs := elastic.NewTermsAggregation().Field("interests")
	res, err = client.Search("lpwtest").Type("employee").Aggregation("all_interests",aggs).Do(ctx)
	fmt.Println("分析： 输出")
	printEmployee(res, err)*/
}
//简单分页
func list(size,page int){
	if size < 0 || page < 1 {
		fmt.Println("param error")
		return
	}
	res, err := client.Search("lpwtest").Type("employee").Size(size).From((page-1)*size).Do(ctx)
	printEmployee(res, err)

}

func main()  {
	fmt.Println("begin CRUD ... ")
	//create()
	//delete()
	//update()
	//get()
	//query()
	list(2,2)
}