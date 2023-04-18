package service

import (
	globalmodel "github.com/GarnBarn/common-go/model"
	"github.com/GarnBarn/gb-tag-consumer/repository"
	"github.com/pquerna/otp/totp"
	"github.com/sirupsen/logrus"
)

type TagService interface {
	CreateTag(tag *globalmodel.Tag) error
	DeleteTag(tagId int) error
}

type tagService struct {
	tagRepository repository.TagRepository
}

func NewTagsService(tagRepository repository.TagRepository) TagService {
	return &tagService{tagRepository: tagRepository}
}

func (t *tagService) CreateTag(tag *globalmodel.Tag) error {

	// Create the otp secret
	totpKeyResult, err := totp.Generate(totp.GenerateOpts{Issuer: "GarnBarn", AccountName: "GarnBarn"})
	if err != nil {
		logrus.Error(err)
		return err
	}
	totpPrivateKey := totpKeyResult.Secret()
	logrus.Info(totpPrivateKey)

	tag.SecretKeyTotp = totpPrivateKey

	return t.tagRepository.Create(tag)
}

func (t *tagService) DeleteTag(tagId int) error {
	logrus.Info("Check delete tag")
	defer logrus.Info("Complete delete tag")
	err := t.tagRepository.Delete(tagId)
	return err
}
