package config

var QueueConfig *RabbitmqConfigType

type RabbitmqConfigType struct {
	AmqpURI string
	Exchange string
	RoutingKey string
}
