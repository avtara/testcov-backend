package service

import (
	"fmt"
	"log"

	"github.com/avtara/testcov-backend/dto"
	"github.com/avtara/testcov-backend/entity"
	"github.com/avtara/testcov-backend/repository"
	"github.com/mashingan/smapping"
)

//OrderService is a contract about something that service can do
type OrderService interface {
	CreateOrder(order dto.OrderDTO) entity.Order
	HistoryOrder(userID int) []entity.Order
}

type orderService struct {
	orderRepository repository.OrderRepository
}

//NewOrderService creates a new instance of AuthService
func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (service *orderService) CreateOrder(order dto.OrderDTO) entity.Order {
	createOrder := entity.Order{}
	err := smapping.FillStruct(&createOrder, smapping.MapFields(&order))
	fmt.Println(err)
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.orderRepository.CreateOrder(createOrder)
	return res
}

func (service *orderService) HistoryOrder(userID int) []entity.Order {
	return service.orderRepository.HistoryOrder(userID)
}
