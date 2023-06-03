package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ybi.com/appbendaharaybi/auth"
	"ybi.com/appbendaharaybi/handler"
	"ybi.com/appbendaharaybi/user"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/appbendaharaybi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	
	userRepository:= user.NewRepository(db)


	userService := user.NewService(userRepository)
	authService := auth.NewService()

	
	userHandler := handler.NewUserHandler(userService, authService)



	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	api.POST("/login", userHandler.Login)



	router.Run(":8080")
}