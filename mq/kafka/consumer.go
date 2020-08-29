package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"time"
)

var consumer *cluster.Consumer

type Handler interface {
	Do(topic string, msg []byte) error
}

type Config struct {
	Host []string
	Group string
	Topics []string
}

func InitConsumer(config *Config) {
	fmt.Println("init kafka consumer, it may take a few seconds...")

	var err error

	clusterCfg := cluster.NewConfig()

	clusterCfg.Consumer.Return.Errors = true
	clusterCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	clusterCfg.Consumer.MaxProcessingTime = time.Second * 3
	clusterCfg.Group.Return.Notifications = true
	clusterCfg.Consumer.Offsets.CommitInterval = time.Second
	//clusterCfg.ClientId = config.Group

	clusterCfg.Version = sarama.V0_10_2_1
	if err = clusterCfg.Validate(); err != nil {
		panic(fmt.Sprintf("Kafka consumer config invalidate. config: %v. err: %v", *clusterCfg, err))
	}
	consumer, err = cluster.NewConsumer(config.Host, config.Group, config.Topics, clusterCfg)
	if err != nil {
		panic(fmt.Sprintf("Create kafka consumer error: %v. config: %v", err, clusterCfg))
	}
	log.Info(fmt.Sprintf("load kafak consumer success conn %v",config.Host ))
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
