package entity

import "time"

//Order represents order table in database
type Order struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Date        time.Time `json:"date"`
	TotalPerson int       `gorm:"type:int(3)" json:"total_person"`
	HospitalID  uint64    `gorm:"not null" json:"-"`
	UserID      uint64    `gorm:"not null" json:"-"`
	User        User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Hospital    Hospital  `gorm:"foreignkey:HospitalID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hospital"`
}
