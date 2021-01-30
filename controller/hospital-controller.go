package controller

import "github.com/gin-gonic/gin"

//HospitalController is a contract about something that service can do
type HospitalController interface {
	All(ctx *gin.Context)
}

type hospitalController struct {
	hospitalService service.HospitalService
}

//NewHospitalController create a new instances of BoookController
func NewHospitalController(hospitalService service.HospitalService) HospitalService {
	return &hospitalService{
		hospitalService: hospitalService,
	}
}
