package repository

import (
	"github.com/avtara/testcov-backend/entity"
	"gorm.io/gorm"
)

//HospitalRepository is contract what userRepository can do to db
type HospitalRepository interface {
	AllHospital() []entity.Hospital
}

type hospitalConnection struct {
	connection *gorm.DB
}

//NewHospitalRepository is creates a new instance of UserRepository
func NewHospitalRepository(db *gorm.DB) HospitalRepository {
	return &hospitalConnection{
		connection: db,
	}
}

func (db *hospitalConnection) AllHospital() []entity.Hospital {
	var hospitals []entity.Hospital
	db.connection.Find(&hospitals)
	return hospitals
}
