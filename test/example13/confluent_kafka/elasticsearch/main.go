package main

import (
	"time"
	"github.com/olivere/elastic"  //最新版本v6
	"context"
	"fmt"
	"reflect"
)

type Tweet struct {
	User 		string					`josn:"user"`
	Message 	string					`json:"message"`
	Retweets 	int						`json:"retweets"`
	Image 		string					`json:"image,omitempty"`
	Created 	time.Time				`json: created,omitempty`
	Tags 		[]string				`json:"tags,omitempty"`
	Location 	string					`json:"location,omitempty"`
	Suggest 	*elastic.SuggestField	`json:"suggest_field,omitempty"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

func main()  {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping("http://10.9.1.213:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("elasticsearch returned with code %d and version %s\n",code,info.Version.Number)
	esversion, err := client.ElasticsearchVersion("http://10.9.1.213:9200")
	if err != nil {
		panic(err)
	}
	fmt.Printf("elasticsearch version %s\n",esversion)
	exists, err := client.IndexExists("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex , err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			fmt.Println("not acknowledged")
		}
	}
	tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	put1, err := client.Index().Index("twitter").Type("tweet").Id("1").BodyJson(tweet1).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("indexed tweet %s to index %s,type %s\n",put1.Index,put1.Type)
	tweet2 :=  `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	put2, err := client.Index().Index("twitter").Type("tweet").Id("2").BodyString(tweet2).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("indexed tweet %s to index %s,type %s\n"),put2.Id,put2.Index,put2.Type

	get1, err := client.Get().Index("tweeter").Type("tweet").Id("1").Do(ctx)
	if err != nil {
		panic(err)
	}
	if get1.Found{
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
	_, err = client.Flush().Index("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index("twitter").   // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}
	// T

}