package etcd

import (
	"context"
	"encoding/json"
	"logagent/taillog"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/kataras/golog"
)

var client *clientv3.Client

// Connect with etcd
func Connect(addrs []string, timeout int) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: time.Duration(timeout) * time.Second,
	})
	if err != nil {
		golog.Error("[etcd] Failed to connect with etcd, err:", err)
		return
	}
	golog.Info("[etcd] Success to connect with etcd.")
	return
}

// GetConfig from etcd by key
func GetConfig(key string) (configs []*taillog.TailTaskConfig, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	responses, err := client.Get(ctx, key)
	cancel()
	if err != nil {
		golog.Error("[etcd] Get configs failed. err: ", err)
		return
	}
	for _, kvs := range responses.Kvs {
		golog.Infof("[etcd] Get key: %s, value: %s", string(kvs.Key), string(kvs.Value))
		err = json.Unmarshal(kvs.Value, &configs)
		if err != nil {
			golog.Error("[etcd] json.Unmarshal failed. err: ", err)
			break
		}
		golog.Debug("[etcd] Get configs: ", debugtool(configs))
	}
	return
}

var debugtool = taillog.Debugtool

// BindingConfigChannel Binding with tail task manager.
func BindingConfigChannel(key string, configChan chan<- []*taillog.TailTaskConfig) {
	watchChan := client.Watch(context.Background(), key)

	for response := range watchChan {
		for _, event := range response.Events {
			golog.Debugf("[etcd] Get response from watchChannel, type:%v, key:%s, value:%s",
				event.Type, string(event.Kv.Key), string(event.Kv.Value))
			var newConfigs []*taillog.TailTaskConfig
			if event.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(event.Kv.Value, &newConfigs)
				if err != nil {
					golog.Error("[etcd] unable to unmarshal data from watchChannel. err: ", err)
					break
				}
			}
			golog.Debug("[etcd] Send new configs to channel: ", debugtool(newConfigs))
			configChan <- newConfigs
		}
	}
}
