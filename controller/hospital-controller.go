package controller

import (
	"math"
	"net/http"
	"sort"
	"strconv"

	"github.com/avtara/testcov-backend/entity"
	"github.com/avtara/testcov-backend/helper"
	"github.com/avtara/testcov-backend/service"
	"github.com/gin-gonic/gin"
)

//HospitalController is a contract about something that service can do
type HospitalController interface {
	All(ctx *gin.Context)
	NearestHospital(ctx *gin.Context)
	DetailHospital(ctx *gin.Context)
}

type hospitalController struct {
	hospitalService service.HospitalService
}

//NewHospitalController create a new instances of BoookController
func NewHospitalController(hospitalService service.HospitalService) HospitalController {
	return &hospitalController{
		hospitalService: hospitalService,
	}
}

func (c *hospitalController) All(ctx *gin.Context) {
	var hospitals []entity.Hospital = c.hospitalService.All()
	for i := 0; i < len(hospitals); i++ {
		hospitals[i].Schedules = c.hospitalService.DetailSchedule(strconv.Itoa(i + 1))
	}
	res := helper.BuildResponse(true, "OK", hospitals)
	ctx.JSON(http.StatusOK, res)
}

func (c *hospitalController) NearestHospital(ctx *gin.Context) {
	var hospitals []entity.Hospital = c.hospitalService.All()
	longitude := ctx.Query("longitude")
	latitude := ctx.Query("latitude")
	if longitude == "" && latitude == "" {
		response := helper.BuildErrorResponse("Failed to process request", "Not find latitude and longitude, latitude and longitude needed!", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	for i := 0; i < len(hospitals); i++ {
		fromLong, _ := strconv.ParseFloat(longitude, 64)
		fromLat, _ := strconv.ParseFloat(latitude, 64)
		toLong, _ := strconv.ParseFloat(hospitals[i].Longitude, 64)
		toLat, _ := strconv.ParseFloat(hospitals[i].Latitude, 64)
		hospitals[i].Distance = c.Distance(fromLong, fromLat, toLong, toLat)
	}
	for i := 0; i < len(hospitals); i++ {
		if hospitals[i].Distance >= 1.5 {
			hospitals = append(hospitals[:i], hospitals[i+1:]...)
			i--
		}
	}
	sort.Slice(hospitals, func(i, j int) bool {
		return hospitals[i].Distance < hospitals[j].Distance
	})

	res := helper.BuildResponse(true, "OK", hospitals)
	ctx.JSON(http.StatusOK, res)
}

func (c *hospitalController) DetailHospital(ctx *gin.Context) {
	hospital := c.hospitalService.DetailSchedule(ctx.Param("id"))
	res := helper.BuildResponse(true, "OK", hospital)
	ctx.JSON(http.StatusOK, res)
}

func (c *hospitalController) Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	dist = dist * 1.609344

	return dist
}
