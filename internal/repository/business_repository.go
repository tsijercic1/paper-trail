package repository

import "paper-trail/internal/model"

type BusinessRepository interface {
	Create(business *model.Business) error
	GetByID(id int) (*model.Business, error)
	Update(business *model.Business) error
	Delete(id int) error
	Page(offset, limit int) ([]*model.Business, error)
}
