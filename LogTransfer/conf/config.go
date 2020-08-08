package conf

// AppConf ...
type AppConf struct {
	KafkaConf `ini:"kafka"`
	ESConf    `ini:"es"`
}

// KafkaConf ...
type KafkaConf struct {
	Address string   `ini:"address"`
	Topics  []string `ini:"topics"`
}

// ESConf ...
type ESConf struct {
	Address  string `ini:"address"`
	ChanSize int    `ini:"chansize"`
	Num      int    `ini:"num"`
}
