TMP_FOLDER=./tmp
KAFKA_FOLDER=./kafka
ETCD_FOLDER=./etcd
wget -c -P ${TMP_FOLDER} https://mirrors.tuna.tsinghua.edu.cn/apache/kafka/2.5.0/kafka_2.13-2.5.0.tgz
mkdir -p ${KAFKA_FOLDER}
tar xzvf ${TMP_FOLDER}/kafka_2.13-2.5.0.tgz -C ${KAFKA_FOLDER} --strip-components=1


# ETCD_VER=v3.4.10
# GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
# DOWNLOAD_URL=${GITHUB_URL}

# # rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
# # rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

# curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o ${TMP_FOLDER}/etcd-${ETCD_VER}-linux-amd64.tar.gz
# mkdir -p ${ETCD_FOLDER}
# tar xzvf ${TMP_FOLDER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -C ${ETCD_FOLDER} --strip-components=1
# # rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

# ${ETCD_FOLDER}/etcd --version
# ${ETCD_FOLDER}/etcdctl version