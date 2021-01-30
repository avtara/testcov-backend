package repository

import (
	"github.com/avtara/testcov-backend/entity"
	"gorm.io/gorm"
)

//HospitalRepository is contract what userRepository can do to db
type HospitalRepository interface {
	AllHospital() []entity.Hospital
	DetailSchedule(hospitalID string) []entity.Schedule
	DetailHospital(hospitalID string) entity.Hospital
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

func (db *hospitalConnection) DetailSchedule(hospitalID string) []entity.Schedule {
	var schedule []entity.Schedule
	db.connection.Table("schedules").Select("schedules.id, schedules.day, schedules.time_start, schedules.time_end, schedules.hospital_id").Joins("left join hospitals on schedules.id = hospitals.id").Where("schedules.hospital_id = ?", hospitalID).Scan(&schedule)
	return schedule
}

func (db *hospitalConnection) DetailHospital(hospitalID string) entity.Hospital {
	var hospitals entity.Hospital
	db.connection.Find(&hospitals, hospitalID)
	return hospitals
}
