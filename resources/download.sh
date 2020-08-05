TMP_FOLDER=./tmp

KAFKA_FOLDER=./kafka
KAFKA_TUNA_URL=https://mirrors.tuna.tsinghua.edu.cn/apache/kafka/2.5.0/kafka_2.13-2.5.0.tgz

ETCD_FOLDER=./etcd
ETCD_VER=v3.4.10
ETCD_GOOGLE_URL=https://storage.googleapis.com/etcd
ETCD_GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
ETCD_DOWNLOAD_URL=${ETCD_GITHUB_URL}

wget -c -P ${TMP_FOLDER} ${KAFKA_TUNA_URL} &

# rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
# rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test
curl -L ${ETCD_DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o ${TMP_FOLDER}/etcd-${ETCD_VER}-linux-amd64.tar.gz

mkdir -p ${ETCD_FOLDER}
tar xzvf ${TMP_FOLDER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -C ${ETCD_FOLDER} --strip-components=1
# rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
mkdir -p ${KAFKA_FOLDER}
tar xzvf ${TMP_FOLDER}/kafka_2.13-2.5.0.tgz -C ${KAFKA_FOLDER} --strip-components=1

${ETCD_FOLDER}/etcd --version
${ETCD_FOLDER}/etcdctl version