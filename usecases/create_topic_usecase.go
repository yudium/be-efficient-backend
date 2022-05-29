package usecases

import "be-efficient/config"

type CreateTopicUsecase struct {
}

func (o *CreateTopicUsecase) Execute() string {
	defaultName := "Untitled"
	id := config.Context.TopicGateway.Create(defaultName)
	return id
}

func NewCreateTopicUsecase() *CreateTopicUsecase {
	return &CreateTopicUsecase{}
}
