package config

var QueueConfig *QueueConfigType

type QueueConfigType struct {
	KafkaURI string
	Exchange string
	RoutingKey string
}
