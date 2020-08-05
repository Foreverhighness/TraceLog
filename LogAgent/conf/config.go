package conf

// AppConf application config
type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

// KafkaConf kafka config
type KafkaConf struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

// EtcdConf etcd config
type EtcdConf struct {
	Address       string `ini:"address"`
	Timeout       int    `ini:"timeout"`
	CollectLogKey string `ini:"collect_log_key"`
}
