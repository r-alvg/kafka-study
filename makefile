KAFKA=kafka
KAFKA_BIN=${KAFKA}/bin
KAFKA_CONF=${KAFKA}/config

kafka-producer:
	${KAFKA_BIN}/kafka-console-producer.sh --broker-list localhost:9092 --topic CARD_PGT

kafka-server:
	${KAFKA_BIN}/kafka-server-start.sh ${KAFKA_CONF}/server.properties

kafka-zookepper:
	${KAFKA_BIN}/zookeeper-server-start.sh ${KAFKA_CONF}/zookeeper.properties

kafka: kafka-zookepper kafka-server