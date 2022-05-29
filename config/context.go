package config

import (
	"be-efficient/gateways"
)

type ContextType struct {
	TopicGateway gateways.TopicGateway
}

var Context ContextType
