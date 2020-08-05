package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/kataras/golog"
)

type logData struct {
	topic, data string
}

var (
	client  sarama.SyncProducer
	logChan chan *logData
)

// Connect connect to kafka
func Connect(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	{
		config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
		config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
		config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	}
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		golog.Error("[kakfa] Fail to connect with kafka. err:", err)
		return
	}
	golog.Info("[kafka] Connect success.")
	logChan = make(chan *logData, maxSize)
	go sendToKafka()
	return
}

// SendToChan topic data
func SendToChan(topic, data string) {
	logChan <- &logData{topic, data}
}

// SendToKafka sent message to kafka by topic
func sendToKafka() {
	msg := &sarama.ProducerMessage{}
	for {
		select {
		case logdata := <-logChan:
			msg.Topic = logdata.topic
			msg.Value = sarama.StringEncoder(logdata.data)
			// 发送消息
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				golog.Error("[kafka] Send message failed, err:", err)
				return
			}
			golog.Infof("[kafka] Success to send, pid: %v offset: %v", pid, offset)
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
