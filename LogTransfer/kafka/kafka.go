package kafka

import (
	"sync"

	"github.com/Shopify/sarama"
	"github.com/kataras/golog"
)

// Worker ...
type Worker struct {
	Do       func(string)
	topic    string
	consumer sarama.Consumer
}

// NewWorker ...
func NewWorker(addrs []string, topic string) (worker *Worker, err error) {
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		golog.Error("[kafka] Fail to create consumer. err: ", err)
		return
	}
	worker = &Worker{
		consumer: consumer,
		topic:    topic,
	}
	return
}

// Run ...
func (w *Worker) Run() {
	var wg sync.WaitGroup
	partitionList, err := w.consumer.Partitions(w.topic)
	if err != nil {
		golog.Error("[kafka] Fail to get list of partition. err: ", err)
		return
	}
	golog.Debug("[kafka] Success to get list of partition: ", partitionList)

	for partition := range partitionList {
		partitionConsumer, err := w.consumer.ConsumePartition(w.topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			golog.Error("[kafka] Fail to create partition consumer. err: ", err)
			return
		}
		golog.Info("[kafka] Success to create partition consumer.")
		defer partitionConsumer.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for message := range partitionConsumer.Messages() {
				golog.Debugf("[kafka] Partition: %d, Offset: %d, key: %v, value: %v.", message.Partition, message.Offset, message.Key, string(message.Value))
				w.Do(string(message.Value))
			}
			wg.Done()
		}(partitionConsumer)
	}
	wg.Wait()
}
