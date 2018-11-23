package confluent_producter

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
)

func main()  {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":"localhost"})
	if err != nil {
		panic(err)
	}
	defer p.Close()
	//delivery report handler for produced messages
	go func(){
		for e := range p.Events(){
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("delivery failed: %v\n",ev.TopicPartition)
				}else {
					fmt.Printf("delivered message to %v\n",ev.TopicPartition)
				}
			}
		}
	}()
	//produce messages to topic(asynchronously)
	topic := "myTopic"
	for _,word := range []string{"welcome","to","the","confluent","kafka","golang","client"}{
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic,Partition: kafka.PartitionAny},
			Value: []byte(word),
		},nil)
	}
	p.Flush(15 * 1000)
}