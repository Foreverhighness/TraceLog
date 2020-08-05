ZOOKEEPER_LOG_FOLDER=./tmp/zookeeper
KAFKA_LOG_FOLDER=./tmp/kafka-logs
KAFKA_FOLDER=./kafka
# zookeeper config file
sed -i "/^dataDir=/c\dataDir=${ZOOKEEPER_LOG_FOLDER}" ${KAFKA_FOLDER}/config/zookeeper.properties
# kafka config file
sed -i "/^log.dirs=/c\log.dirs=${KAFKA_LOG_FOLDER}" ${KAFKA_FOLDER}/config/server.properties

