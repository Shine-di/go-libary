package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"go-libary/log"
	"time"
)

var consumer *cluster.Consumer

type Handler interface {
	Do(topic string, msg []byte) error
}

func InitKafkaConsumer(serverName string) {
	fmt.Println("init kafka consumer, it may take a few seconds...")

	var err error

	clusterCfg := cluster.NewConfig()

	clusterCfg.Consumer.Return.Errors = true
	clusterCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	clusterCfg.Consumer.MaxProcessingTime = time.Second * 3
	clusterCfg.Group.Return.Notifications = true
	clusterCfg.Consumer.Offsets.CommitInterval = time.Second
	clusterCfg.ClientId = serverName

	clusterCfg.Version = sarama.V0_10_2_1
	if err = clusterCfg.Validate(); err != nil {
		msg := fmt.Sprintf("Kafka consumer config invalidate. config: %v. err: %v", *clusterCfg, err)
		fmt.Println(msg)
		panic(msg)
	}

	consumer, err = cluster.NewConsumer(brokerList, serverName, topics, clusterCfg)
	if err != nil {
		msg := fmt.Sprintf("Create kafka consumer error: %v. config: %v", err, clusterCfg)
		fmt.Println(msg)
		log.Fatal(msg)
	}

}

func Stop() {
	consumer.Close()
}

func Subscribe(handlers Handler) {

	for {
		select {
		case msg, more := <-consumer.Messages():
			if more {
				kafkaConsumer(msg, handlers)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case err, more := <-consumer.Errors():
			if more {
				fmt.Println("Kafka consumer error: %v", err.Error())
				log.Fatal(err.Error())
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				fmt.Println("Kafka consumer rebalance: %v", ntf)
			}
		}
	}
}

func kafkaConsumer(msg *sarama.ConsumerMessage, handler Handler) {
	if err := handler.Do(msg.Topic, msg.Value); err != nil {
		log.Error(err.Error())
	}
}
