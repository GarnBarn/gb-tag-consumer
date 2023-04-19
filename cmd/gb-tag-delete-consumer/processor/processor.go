package processor

import (
	"encoding/json"
	globalmodel "github.com/GarnBarn/common-go/model"
	rabbitMQ "github.com/GarnBarn/common-go/rabbitmq"
	"github.com/GarnBarn/gb-tag-consumer/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type Processor struct {
	rabbitmqPublisher *rabbitmq.Publisher
	tagRepository     repository.TagRepository
}

func NewProcessor(rabbitmqPublisher *rabbitmq.Publisher, tagRepository repository.TagRepository) rabbitMQ.Processor {
	return &Processor{
		rabbitmqPublisher: rabbitmqPublisher,
		tagRepository:     tagRepository,
	}
}

func (p *Processor) Process(d rabbitmq.Delivery) error {
	tagRequest := globalmodel.TagDeleteRequest{}
	err := json.Unmarshal(d.Body, &tagRequest)
	if err != nil {
		logrus.Error("Can't unmarshal data: ", err)
		return err
	}
	err = p.tagRepository.Delete(tagRequest.ID)
	if err != nil {
		logrus.Error("Can't delete data: ", err)
		return err
	}

	logrus.Info("Successfully deleted the tag id: ", tagRequest.ID)
	return nil
}
