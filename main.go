package main

import (
	"be-efficient/config"
	"be-efficient/gateways"
)

func main() {
	config.Context = config.ContextType{
		TopicGateway: gateways.NewTopicGateway(),
	}
}
