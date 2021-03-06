package usecases_test

import (
	"be-efficient/config"
	"be-efficient/usecases"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type CreateMetadata struct {
	IsCalled  bool
	NameParam string
}

type TopicGatewaySpy struct {
	CreateMetadata CreateMetadata
}

var idFromCreateMethod = "gsau2y8f"

func (o *TopicGatewaySpy) Create(name string) string {
	o.CreateMetadata.IsCalled = true
	o.CreateMetadata.NameParam = name
	return idFromCreateMethod
}

var _ = Describe("Create Topic", func() {
	var topicGatewaySpy *TopicGatewaySpy
	var usecase *usecases.CreateTopicUsecase

	BeforeEach(func() {
		usecase = usecases.NewCreateTopicUsecase()
		topicGatewaySpy = &TopicGatewaySpy{}
		config.Context = config.ContextType{
			TopicGateway: topicGatewaySpy,
		}
	})

	It("calls repo for creating topic", func() {
		usecase.Execute()
		Expect(topicGatewaySpy.CreateMetadata.IsCalled).To(Equal(true))
	})

	It("create topic with default name", func() {
		defaultName := "Untitled"
		usecase.Execute()
		Expect(topicGatewaySpy.CreateMetadata.NameParam).To(Equal(defaultName))
	})

	It("returns topic id", func() {
		Expect(usecase.Execute()).To(Equal(idFromCreateMethod))
	})
})
