package service

import (
	"github.com/avtara/testcov-backend/entity"
	"github.com/avtara/testcov-backend/repository"
)

//HospitalService is a contract about something that service can do
type HospitalService interface {
	All() []entity.Hospital
	DetailSchedule(hospitalID string) []entity.Schedule
	DetailHospital(hospitalID string) entity.Hospital
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

func (service *hospitalService) DetailSchedule(hospitalID string) []entity.Schedule {
	return service.hospitalRepository.DetailSchedule(hospitalID)
}

func (service *hospitalService) DetailHospital(hospitalID string) entity.Hospital {
	return service.hospitalRepository.DetailHospital(hospitalID)
}
