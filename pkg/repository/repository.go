package repository

import (
	"github.com/GarnBarn/common-go/model"
	"gorm.io/gorm"
)

type TagRepository interface {
	Create(tag *model.Tag) error
	Delete(tagId int) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	db.AutoMigrate(&model.Tag{})
	return &tagRepository{
		db: db,
	}
}

func (t *tagRepository) Create(tag *model.Tag) error {
	result := t.db.Create(tag)
	return result.Error
}

func (t *tagRepository) Delete(tagId int) error {
	res := t.db.Delete(&model.Tag{}, tagId)
	return res.Error
}
