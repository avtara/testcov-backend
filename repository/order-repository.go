package repository

import (
	"github.com/avtara/testcov-backend/entity"
	"gorm.io/gorm"
)

//OrderRepository is contract what userRepository can do to db
type OrderRepository interface {
	CreateOrder(order entity.Order) entity.Order
	HistoryOrder(userID int) []entity.Order
}

type orderConnection struct {
	connection *gorm.DB
}

//NewOrderRepository is creates a new instance of UserRepository
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderConnection{
		connection: db,
	}
}

func (db *orderConnection) CreateOrder(order entity.Order) entity.Order {
	db.connection.Save(&order)
	return order
}

func (db *orderConnection) HistoryOrder(userID int) []entity.Order {
	var order []entity.Order
	db.connection.Where("user_id = ?", userID).Find(&order)
	return order
}
