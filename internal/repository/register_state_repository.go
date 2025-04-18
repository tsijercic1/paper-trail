package repository

import "paper-trail/internal/model"

type RegisterStateRepository interface {
	Create(state *model.RegisterState) error
	GetByID(id int) (*model.RegisterState, error)
	Update(state *model.RegisterState) error
	Delete(id int) error
	List() ([]*model.RegisterState, error)
}
