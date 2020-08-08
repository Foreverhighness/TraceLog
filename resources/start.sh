KAFKA_FOLDER=./kafka
ETCD_FOLDER=./etcd
ELASTICSEARCH_FOLDER=./elasticsearch
KIBANA_FOLDER=./kibana
# Start Zookeeper
${KAFKA_FOLDER}/bin/zookeeper-server-start.sh ${KAFKA_FOLDER}/config/zookeeper.properties &
# Start Kafka
${KAFKA_FOLDER}/bin/kafka-server-start.sh ${KAFKA_FOLDER}/config/server.properties &
# Start Etcd
${ETCD_FOLDER}/etcd &
# Start Elasticsearch
${ELASTICSEARCH_FOLDER}/bin/elasticsearch &
# Start Kibana
${KIBANA_FOLDER}/bin/kibana &

