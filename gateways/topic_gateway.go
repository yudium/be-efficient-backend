package gateways

type TopicGateway interface {
	Create(name string) string
}

type TopicGatewayMySQL struct {
}

func (o *TopicGatewayMySQL) Create(name string) string {
	return "REPLACE THIS"
}

func NewTopicGateway() TopicGateway {
	return &TopicGatewayMySQL{}
}
