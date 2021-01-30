package main

import (
	"github.com/avtara/testcov-backend/config"
	"github.com/avtara/testcov-backend/controller"
	"github.com/avtara/testcov-backend/middleware"
	"github.com/avtara/testcov-backend/repository"
	"github.com/avtara/testcov-backend/service"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                      = config.SetupDatabaseConnection()
	userRepository     repository.UserRepository     = repository.NewUserRepository(db)
	hospitalRepository repository.HospitalRepository = repository.NewHospitalRepository(db)
	orderRepository    repository.OrderRepository    = repository.NewOrderRepository(db)
	jwtService         service.JWTService            = service.NewJWTService()
	authService        service.AuthService           = service.NewAuthService(userRepository)
	authController     controller.AuthController     = controller.NewAuthController(authService, jwtService)
	hospitalService    service.HospitalService       = service.NewHospitalService(hospitalRepository)
	hospitalController controller.HospitalController = controller.NewHospitalController(hospitalService)
	orderService       service.OrderService          = service.NewOrderService(orderRepository)
	orderController    controller.OrderController    = controller.NewOrderController(orderService, jwtService, hospitalService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	config := cors.DefaultConfig()
	r.Use(cors.Default())
	// r.Use(cors.Middleware(cors.Config{
	// 	Origins:         "*",
	// 	Methods:         "GET, PUT, POST, DELETE",
	// 	RequestHeaders:  "Origin, Authorization, Content-Type",
	// 	ExposedHeaders:  "",
	// 	MaxAge:          50 * time.Second,
	// 	Credentials:     true,
	// 	ValidateHeaders: false,
	// }))

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.GET("/validate", authController.ValidateToken, middleware.AuthorizeJWT(jwtService))
	}

	hospitalRoutes := r.Group("api/hospital")
	{
		hospitalRoutes.GET("/", hospitalController.All)
		hospitalRoutes.GET("/nearest", hospitalController.NearestHospital)
		hospitalRoutes.GET("/detail/:id", hospitalController.DetailHospital)
	}

	orderRoutes := r.Group("api/order")
	{
		orderRoutes.POST("/add", orderController.CreateOrder)
		orderRoutes.GET("/", orderController.HistoryOrder, middleware.AuthorizeJWT(jwtService))
	}

	r.Run()
}
