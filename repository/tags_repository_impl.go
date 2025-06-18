package repository

import (
	"CRUD/data/request"
	"CRUD/helper"
	"CRUD/model"
	"errors"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t *TagsRepositoryImpl) FindAll() (tags []model.Tags) {
	var tag []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tag
}

// Delete implements TagsRepository.
func (t *TagsRepositoryImpl) Delete(tagsID int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsID).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository.
func (t *TagsRepositoryImpl) FindById(tagsId int) (model.Tags, error) {
	var tag model.Tags
	result := t.Db.First(&tag, tagsId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return tag, errors.New("tag not found")
	}
	return tag, result.Error
}

// Save implements TagsRepository.
func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

// Update implements TagsRepository.
func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	updateTag := request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}
