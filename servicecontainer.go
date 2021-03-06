package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/opam22/burgers/controllers"
	"github.com/opam22/burgers/infrastructures"
	"github.com/opam22/burgers/repositories"
	"github.com/opam22/burgers/services"
)

type IServiceContainer interface {
	InjectMenuController() controllers.MenuController
	InjectOrderController() controllers.OrderController
}

type kernel struct{}

func (k *kernel) InjectMenuController() controllers.MenuController {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sqlConn, _ := sql.Open("mysql", os.Getenv("MYSQL_HOST"))
	DBHandler := &infrastructures.MySQLHandler{}
	DBHandler.Conn = sqlConn

	menuRepository := &repositories.MenuRepository{DBHandler}
	menuService := &services.MenuService{menuRepository}
	menuController := controllers.MenuController{menuService}

	return menuController
}

func (k *kernel) InjectOrderController() controllers.OrderController {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sqlConn, _ := sql.Open("mysql", os.Getenv("MYSQL_HOST"))
	DBHandler := &infrastructures.MySQLHandler{}
	DBHandler.Conn = sqlConn

	orderRepository := &repositories.OrderRepository{DBHandler}
	orderService := &services.OrderService{orderRepository}
	orderController := controllers.OrderController{orderService}

	return orderController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
