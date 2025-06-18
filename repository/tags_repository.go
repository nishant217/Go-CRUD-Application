package repository

import "CRUD/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsID int)
	FindById(tagsID int) (tags model.Tags, err error)
	FindAll() (tags []model.Tags)
}
