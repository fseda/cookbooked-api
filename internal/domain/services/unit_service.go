package services

import (
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/repositories"
)

type UnitService interface {
	GetAllUnits() ([]models.Unit, error)
}

type unitService struct {
	unitRepository repositories.UnitRepository
}

func NewUnitService(unitRepository repositories.UnitRepository) UnitService {
	return &unitService{unitRepository}
}

func (s *unitService) GetAllUnits() ([]models.Unit, error) {
	return s.unitRepository.GetAllUnits()
}