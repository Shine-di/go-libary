package kafka

import (
	"context"
	"errors"
	"fmt"
	"go-libary/log"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

const PusherMsgLength = 1000

var producerCli sarama.AsyncProducer
var pusherMsg = make(chan *sarama.ProducerMessage, PusherMsgLength)

func Push(ctx context.Context, key, topic string, msg string) error {
	if producerCli == nil {
		return errors.New("请先设置MQProducer为True! ")
	}
	if len(topic) == 0 {
		return errors.New("Topic为空")
	}
	if len(msg) == 0 {
		return errors.New("Msg为空")
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	if len(pusherMsg) == PusherMsgLength {
		return errors.New("消息队列阻塞")
	}
	pusherMsg <- message
	return nil
}

func InitProducer(config *Config) {
	configPro := sarama.NewConfig()
	configPro.Producer.Retry.Max = 5
	configPro.Producer.RequiredAcks = sarama.WaitForAll
	configPro.Producer.Return.Successes = true
	configPro.Producer.Partitioner = sarama.NewManualPartitioner
	producer, err := sarama.NewAsyncProducer(config.Host, configPro)
	if err != nil {
		panic(fmt.Sprintf("init producer fail %v", zap.Error(err)))
	}
	producerCli = producer
	go callback()
	go pushDataToKafka()
	log.Info(fmt.Sprintf("load kafak producer success conn %v",config.Host ))
}

func pushDataToKafka() {
	for {
		select {
		case msg := <-pusherMsg:
			{
				log.Info("push to kafka", zap.Any("msg", msg))
				producerCli.Input() <- msg
			}
		}
	}
}

func callback() {
	for {
		select {
		case success := <-producerCli.Successes():
			log.Info("push to kafka success", zap.Any("msg", success))
		case errInfo := <-producerCli.Errors():
			log.Error("push to kafka error", zap.Error(errInfo))
		}
	}
}
