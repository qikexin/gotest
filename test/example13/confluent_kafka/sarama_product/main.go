package main

import (
	"github.com/Shopify/sarama"
	"fmt"
)

func main()  {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test,my message is good asdfasdfa s")

	client, err := sarama.NewSyncProducer([]string{"10.9.1.213:9092"},config)
	if err != nil {
		fmt.Println("producer close, err: ",err)
		return
	}
	defer client.Close()
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,",err)
		return
	}
	fmt.Printf("pid: %v offset: %v\n",pid,offset)
}