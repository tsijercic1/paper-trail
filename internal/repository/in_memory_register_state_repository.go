package repository

import (
	"errors"
	"paper-trail/internal/model"
)

type InMemoryRegisterStateRepository struct {
	store map[int]*model.RegisterState
}

func NewInMemoryRegisterStateRepository() *InMemoryRegisterStateRepository {
	return &InMemoryRegisterStateRepository{
		store: make(map[int]*model.RegisterState),
	}
}

func (r *InMemoryRegisterStateRepository) Create(state *model.RegisterState) error {
	if _, exists := r.store[state.ID]; exists {
		return errors.New("register state already exists")
	}
	r.store[state.ID] = state
	return nil
}

func (r *InMemoryRegisterStateRepository) GetByID(id int) (*model.RegisterState, error) {
	state, exists := r.store[id]
	if !exists {
		return nil, errors.New("register state not found")
	}
	return state, nil
}

func (r *InMemoryRegisterStateRepository) Update(state *model.RegisterState) error {
	if _, exists := r.store[state.ID]; !exists {
		return errors.New("register state not found")
	}
	r.store[state.ID] = state
	return nil
}

func (r *InMemoryRegisterStateRepository) Delete(id int) error {
	if _, exists := r.store[id]; !exists {
		return errors.New("register state not found")
	}
	delete(r.store, id)
	return nil
}

func (r *InMemoryRegisterStateRepository) List() ([]*model.RegisterState, error) {
	var states []*model.RegisterState
	for _, state := range r.store {
		states = append(states, state)
	}
	return states, nil
}
