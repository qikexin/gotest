package sarama_consumer

import (
	"github.com/Shopify/sarama"
	"strings"
	"fmt"
	"time"
)

func main()  {
	consumer, err := sarama.NewConsumer(strings.Split("120.25.160.52:9092",","),nil))
	if err != nil {
		fmt.Printf("failed to start consumer: %s",err)
		return
	}
	partitionList, err := consumer.Partitions("nginx_log")
	if err != nil {
		fmt.Println("failed to get the list of partitions: ",err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList{
		pc, err := consumer.ConsumePartition("nginx_log",int32(partition),sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d: %s\n",partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages(){
				fmt.Printf("partition: %d,offset: %d, key: %s,value %s\n", msg.Partition,msg.Offset,string(msg.Key))
			}
		}(pc)
		time.Sleep(time.Hour)
		consumer.Close()
	}
}
