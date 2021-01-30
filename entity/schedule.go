package entity

import "time"

//Schedule represents users table in database
type Schedule struct {
	ID         uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Day        string    `gorm:"type:varchar(20)" json:"day"`
	TimeStart  time.Time `gorm:"type:varchar(20)" json:"time_start"`
	TimeEnd    time.Time `gorm:"type:varchar(20)" json:"time_end"`
	HospitalID uint64    `gorm:"not null" json:"-"`
	Hospital   Hospital  `gorm:"foreignkey:HospitalID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hospital"`
}
