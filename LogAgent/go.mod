module logagent

go 1.14

replace (
	github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5
	go.etcd.io/bbolt v1.3.5 => github.com/coreos/bbolt v1.3.5
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/Shopify/sarama v1.26.4
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.6 // indirect
	github.com/hpcloud/tail v1.0.0
	github.com/jonboulle/clockwork v0.2.0 // indirect
	github.com/kataras/golog v0.0.18
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/genproto v0.0.0-20200731012542-8145dea6a485 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/ini.v1 v1.57.0
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
