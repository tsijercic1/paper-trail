package service

import (
	"paper-trail/internal/model"
	"paper-trail/internal/repository"
)

type BusinessService struct {
	repo repository.BusinessRepository
}

func NewBusinessService(repo repository.BusinessRepository) *BusinessService {
	return &BusinessService{repo: repo}
}

func (s *BusinessService) CreateBusiness(business *model.Business) error {
	return s.repo.Create(business)
}

func (s *BusinessService) GetBusinessByID(id int) (*model.Business, error) {
	return s.repo.GetByID(id)
}

func (s *BusinessService) UpdateBusiness(business *model.Business) error {
	return s.repo.Update(business)
}

func (s *BusinessService) DeleteBusiness(id int) error {
	return s.repo.Delete(id)
}

func (s *BusinessService) GetBusinessesPage(offset, limit int) ([]*model.Business, error) {
	return s.repo.Page(offset, limit)
}
