package processor

import (
	"encoding/json"
	globalmodel "github.com/GarnBarn/common-go/model"
	rabbitMQ "github.com/GarnBarn/common-go/rabbitmq"
	"github.com/GarnBarn/gb-tag-consumer/pkg/repository"
	"github.com/pquerna/otp/totp"
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
	tag := globalmodel.Tag{}
	err := json.Unmarshal(d.Body, &tag)
	if err != nil {
		logrus.Error("Can't unmarshal data: ", err)
		return err
	}

	// Create the otp secret
	totpKeyResult, err := totp.Generate(totp.GenerateOpts{Issuer: "GarnBarn", AccountName: "GarnBarn"})
	if err != nil {
		logrus.Error(err)
		return err
	}
	totpPrivateKey := totpKeyResult.Secret()
	logrus.Info(totpPrivateKey)

	tag.SecretKeyTotp = totpPrivateKey

	err = p.tagRepository.Create(&tag)
	if err != nil {
		logrus.Error("Can't save data: ", err)
		return err
	}

	logrus.Info("Successfully created the tag id: ", tag.ID)
	return nil
}
