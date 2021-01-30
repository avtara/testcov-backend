package dto

import "time"

//OrderDTO is a model that used by client when POST from /login url
type OrderDTO struct {
	Date        time.Time `json:"date" form:"date" binding:"required"`
	TotalPerson int       `json:"total_person" form:"total_person" binding:"required"`
	UserID      uint64    `json:"user_id" form:"user_id" binding:"required"`
	HospitalID  uint64    `json:"hospital_id" form:"hospital_id" binding:"required"`
}
