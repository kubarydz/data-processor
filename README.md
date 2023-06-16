# data-processor

Simple data processor, with random data generator

To run locally:
1. Run Kafka:
```
bin/zookeeper-server-start.sh config/zookeeper.properties
bin/kafka-server-start.sh config/server.properties
```

2. Run MongoDB:
```
systemctl start mongod.service
```

3. Run data generator:
```
go run cmd/generator/main.go
```

5. Run data ingester:
```
go run cmd/app/main.go
```
