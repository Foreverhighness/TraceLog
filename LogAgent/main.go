package main

import (
	"logagent/conf"
	"logagent/kafka"
	"logagent/taillog"
	"os"
	"os/signal"
	"syscall"

	"github.com/kataras/golog"
	"gopkg.in/ini.v1"
)

const (
	configPath = "./conf/config.ini"
)

var configs = []*taillog.TailTaskConfig{
	{Path: "./tmp/1.log", Topic: "Mytest_1"},
	{Path: "./tmp/2.log", Topic: "Mytest_2"},
	{Path: "./tmp/3.log", Topic: "Mytest_3"},
}

func init() {
	golog.SetLevel("debug")
}
func wait() {
	exitc := make(chan os.Signal, 1)
	signal.Notify(exitc, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-exitc
}
func main() {
	// Load local config file
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

	// // Connect to ectd
	// etcd.Connect([]string{appConf.EtcdConf.Address}, appConf.EtcdConf.Timeout)
	// // Get remote tail configs
	// configs, err := etcd.GetConfig(appConf.EtcdConf.CollectLogKey)
	// if err != nil {
	// 	golog.Error("[main] Get configs failed. err: ", err)
	// }

	// Start taillog process
	err = taillog.Init(configs)
	if err != nil {
		golog.Error("[main] Fail to start taillog. err:", err)
		return
	}
	golog.Info("[main] Start taillog success")

	// go etcd.BindingConfigChannel(appConf.EtcdConf.CollectLogKey, taillog.GetUpdateChan())

	wait()
}
