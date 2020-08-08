package main

import (
	"fmt"
	"logtransfer/conf"
	"logtransfer/es"
	"os"
	"os/signal"
	"syscall"

	"github.com/kataras/golog"
	"gopkg.in/ini.v1"
)

const (
	configPath = "./conf/config.ini"
)

func wait() {
	exitc := make(chan os.Signal, 1)
	signal.Notify(exitc, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-exitc
}
func init() {
	golog.SetLevel("debug")
}

func factory(topic string) func(string) {
	return func(value string) {
		parameter := &es.Parameter{
			Index: topic,
			Data:  value,
		}
		es.SendToChan(parameter)
		return
	}
}

func main() {
	// Load local config
	var appConfig conf.AppConf
	err := ini.MapTo(&appConfig, configPath)
	if err != nil {
		golog.Error("[main] Failed to open config file. err ", err)
		return
	}
	golog.Info("[main] Success to open config file.")
	// Connect ElasticSearch
	err = es.Init(appConfig.ESConf.Address, appConfig.ESConf.ChanSize)
	if err != nil {
		golog.Error("[main] Fail to connect ElasticSearch. err: ", err)
		return
	}
	es.SendToChan(&es.Parameter{
		Index: "Mytest_1",
		Data:  "This is a test.",
	})
	// Connect kafka
	// workers := make([]*kafka.Worker, 0, len(appConfig.Topics))
	// for _, topic := range appConfig.Topics {
	// 	worker, err := kafka.NewWorker([]string{appConfig.KafkaConf.Address}, topic)
	// 	worker.Do = factory(topic)
	// 	if err != nil {
	// 		golog.Errorf("[main] Fail to creater worker<%s, %s>. err: ", appConfig.KafkaConf.Address, topic, err)
	// 		continue
	// 	}
	// 	golog.Infof("[main] Success to creater worker<%s, %s>.", appConfig.KafkaConf.Address, topic)
	// 	workers = append(workers, worker)
	// }
	// for _, w := range workers {
	// 	go w.Run()
	// }
	wait()
}

func testConfig(appConfig *conf.AppConf) {
	fmt.Println(appConfig)
	fmt.Println(appConfig.KafkaConf)
	fmt.Println(appConfig.ESConf)
	fmt.Println(appConfig.Topics)
	fmt.Println(appConfig.Topics[0])
	fmt.Println(appConfig.Topics[1])
	fmt.Println(appConfig.Topics[2])
}
