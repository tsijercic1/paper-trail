package repository

import (
	"errors"
	"math/rand"
	"paper-trail/internal/model"
)

type InMemoryBusinessRepository struct {
	store map[int]*model.Business
}

func NewInMemoryBusinessRepository() *InMemoryBusinessRepository {
	return &InMemoryBusinessRepository{
		store: make(map[int]*model.Business),
	}
}

func (r *InMemoryBusinessRepository) Create(business *model.Business) error {
	if business.ID == 0 {
		business.ID = rand.Intn(1000000) // Generate a random ID if not provided
	}
	if _, exists := r.store[business.ID]; exists {
		return errors.New("business already exists")
	}
	r.store[business.ID] = business
	return nil
}

func (r *InMemoryBusinessRepository) GetByID(id int) (*model.Business, error) {
	business, exists := r.store[id]
	if !exists {
		return nil, errors.New("business not found")
	}
	return business, nil
}

func (r *InMemoryBusinessRepository) Update(business *model.Business) error {
	if _, exists := r.store[business.ID]; !exists {
		return errors.New("business not found")
	}
	r.store[business.ID] = business
	return nil
}

func (r *InMemoryBusinessRepository) Delete(id int) error {
	if _, exists := r.store[id]; !exists {
		return errors.New("business not found")
	}
	delete(r.store, id)
	return nil
}

func (r *InMemoryBusinessRepository) Page(offset, limit int) ([]*model.Business, error) {
	var businesses []*model.Business
	count := 0
	for _, business := range r.store {
		if count >= offset && count < offset+limit {
			businesses = append(businesses, business)
		}
		count++
	}
	return businesses, nil
}
