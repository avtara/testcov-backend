package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/avtara/testcov-backend/dto"
	"github.com/avtara/testcov-backend/helper"
	"github.com/avtara/testcov-backend/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//OrderController interface is a contract what controller can do
type OrderController interface {
	CreateOrder(ctx *gin.Context)
	HistoryOrder(ctx *gin.Context)
}

type orderController struct {
	orderService    service.OrderService
	jwtService      service.JWTService
	hospitalService service.HospitalService
}

//NewOrderController creates a new instance of AuthController
func NewOrderController(orderService service.OrderService, jwtService service.JWTService, hospitalService service.HospitalService) OrderController {
	return &orderController{
		orderService:    orderService,
		jwtService:      jwtService,
		hospitalService: hospitalService,
	}
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var orderDTO dto.OrderDTO
	errDTO := ctx.ShouldBind(&orderDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	orderResult := c.orderService.CreateOrder(orderDTO)
	response := helper.BuildResponse(true, "OK!", orderResult)
	ctx.JSON(http.StatusOK, response)
	return
}

func (c *orderController) HistoryOrder(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	splitToken := strings.Split(authHeader, "Bearer ")
	token, err := c.jwtService.ValidateToken(splitToken[1])
	if !token.Valid {
		response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	} else {
		userID := c.getUserIDByToken(splitToken[1])
		convertedUserID, err := strconv.Atoi(userID)
		if err == nil {
			orderResult := c.orderService.HistoryOrder(convertedUserID)
			for i := 0; i < len(orderResult); i++ {
				// orderResult[i].User = c.Distance(fromLong, fromLat, toLong, toLat)
				orderResult[i].Hospital = c.hospitalService.DetailHospital(strconv.FormatUint(orderResult[i].HospitalID, 10))
			}
			response := helper.BuildResponse(true, "OK!", orderResult)
			ctx.JSON(http.StatusOK, response)
			return
		}
	}
}

func (c *orderController) getUserIDByToken(token string) string {
	fmt.Println(token)
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
