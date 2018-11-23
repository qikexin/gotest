package kafka

import (
	"github.com/Shopify/sarama"
	"fmt"
	"time"
)

func main()  {
	//设置配置，生产和消费的过程都是通过一个配置开始的
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机的分区类型
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应，只有上面的requireAcks设置不是NoResponse时才有用
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本
	//config.Version = sarama.V0_11_0_0

	client, err := sarama.NewSyncProducer([]string{"192.168.14.4:9092"},config)
	if err != nil {
		fmt.Println("produce close ,err : ",err)
		return
	}
	defer client.Close()

	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "nginx_log"
		msg.Value = sarama.StringEncoder("this is a good test,my message is good")

		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed, ",err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n",pid,offset)
		time.Sleep(10*time.Millisecond)
	}
}