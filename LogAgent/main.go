package main

import (
	"logagent/conf"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"sync"

	"github.com/kataras/golog"
	"gopkg.in/ini.v1"
)

const (
	configPath = "./conf/config.ini"
)

func init() {
	golog.SetLevel("debug")
}
func main() {
	// Load config file
	var appConf conf.AppConf
	err := ini.MapTo(&appConf, configPath)
	if err != nil {
		golog.Error("[main] Fail to load config. err:", err)
		return
	}
	golog.Info("[main] Load config success.")
	golog.Debug("[main] Get config: ", appConf)

	// Connect to kafka
	err = kafka.Connect([]string{appConf.KafkaConf.Address}, appConf.KafkaConf.ChanMaxSize)
	if err != nil {
		golog.Error("[main] Fail to connect kafka. err:", err)
		return
	}
	golog.Info("[main] Connect kafka sucess.")

	// Connect to ectd
	etcd.Connect([]string{appConf.EtcdConf.Address}, appConf.EtcdConf.Timeout)

	configs, err := etcd.GetConfig(appConf.EtcdConf.CollectLogKey)
	if err != nil {
		golog.Error("[main] Get configs failed. err: ", err)
	}

	// Start taillog process
	err = taillog.Init(configs)
	if err != nil {
		golog.Error("[main] Fail to start taillog. err:", err)
		return
	}
	golog.Info("[main] Start taillog success")

	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.BindingConfigChannel(appConf.EtcdConf.CollectLogKey, taillog.GetUpdateChan())
	wg.Wait()
}
