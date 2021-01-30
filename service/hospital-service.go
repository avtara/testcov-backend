package service

import (
	"github.com/avtara/golang_api/entity"
	"github.com/avtara/testcov-backend/repository"
)

//HospitalService is a contract about something that service can do
type HospitalService interface {
	All() []entity.Hospital
}

type hospitalService struct {
	hospitalRepository repository.HospitalRepository
}

//NewHospitalService creates a new instance of AuthService
func NewHospitalService(hospitalRepository repository.HospitalRepository) HospitalService {
	return &hospitalService{
		hospitalRepository: hospitalRepository,
	}
}

func (service *hospitalService) All() []entity.Hospital {
	return service.hospitalRepository.AllHospital()
}
