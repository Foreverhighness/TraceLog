# 查看有多少个 kafka borker 连上了 zookeeper
docker exec -t kafka_zookeeper_1 ./bin/zkCli.sh ls /brokers/ids
# 查看 kafka 里的 topic
docker exec -t kafka_kafka_1 kafka-topics.sh --bootstrap-server :9092 --list
# 在 kafka 里创建一个 topic, 名为 test1
docker exec -t kafka_kafka_1 kafka-topics.sh --bootstrap-server :9092 --create --replication-factor 1 --partitions 1 --topic test1
# 在 kafka 里从头消费 test1
docker exec -t kafka_kafka_1 kafka-console-consumer.sh --bootstrap-server :9092 --topic test1 --from-beginning
# 在 kafka 中查看 test1
docker exec -t kafka_kafka_1 kafka-topics.sh --bootstrap-server :9092 --describe --topic test1
# 在 kafka 中创建一个生产者
docker exec -it kafka_kafka_1 kafka-console-producer.sh  --broker-list :9092 --topic test1
# 往 etcd 里插入键值对，值中包括三个配置
docker exec -t etcd_etcd_1 etcdctl --endpoints=etcd:2379 put /logagnet/collect_config "[{\"path\":\"/data/logs/1.log\",\"topic\":\"test1\"},{\"path\":\"/data/logs/2.log\",\"topic\":\"test2\"},{\"path\":\"/data/logs/3.log\",\"topic\":\"test3\"}]"
# 往 etcd 里插入键值对，值中包括两个配置
docker exec -t etcd_etcd_1 etcdctl --endpoints=etcd:2379 put /logagnet/collect_config "[{\"path\":\"/data/logs/1.log\",\"topic\":\"test1\"},{\"path\":\"/data/logs/2.log\",\"topic\":\"test2\"}]"
# 往 etcd 里查看键值对
docker exec -t etcd_etcd_1 etcdctl --endpoints=etcd:2379 get /logagnet/collect_config

# 以下为试错时的脚本，留作纪念
# kafka-topics.sh --zookeeper zookeeper:2181 --create --replication-factor 1 --partitions 1 --topic test
# kafka-console-consumer.sh --zookeeper zookeeper:2181 --topic test --from-beginning
# docker exec -t kafka_kafka_1 kafka-topics.sh --zookeeper zookeeper:2181 --list
# kafka-console-consumer.sh --bootstrap-server :9092 --topic test --from-beginning
# docker exec -t kafka_kafka_1 kafka-topics.sh --zookeeper zookeeper:2181 --describe --topic test1
# docker exec -t kafka_kafka_1 kafka-topics.sh  -zookeeper zookeeper:2181 --create --topic t1 --partitions 3 --replication-factor 1
# docker exec -t kafka_kafka_1 kafka-topics.sh --zookeeper zookeeper:2181 --create --replication-factor 1 --partitions 1 --topic test2
