KAFKA_FOLDER=./kafka
ETCD_FOLDER=./etcd
# Start Zookeeper
${KAFKA_FOLDER}/bin/zookeeper-server-start.sh ${KAFKA_FOLDER}/config/zookeeper.properties
# Start Kafka
${KAFKA_FOLDER}/bin/kafka-server-start.sh ${KAFKA_FOLDER}/config/server.properties
# Start Etcd
${ETCD_FOLDER}/etcd
