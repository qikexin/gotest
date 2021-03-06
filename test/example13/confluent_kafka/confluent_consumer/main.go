package confluent_consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
)

func main()  {
	c,err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id": "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil{
		panic(err)
	}
	c.SubscribeTopics([]string{"myTopic","^aRegex.*[Tt]opic"},nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("message on %s: %s\n",msg.TopicPartition,string(msg.Value))
		}else {
			fmt.Printf("consumer error: %v (%v)\n",err,msg)
		}
	}
	c.Close()
}