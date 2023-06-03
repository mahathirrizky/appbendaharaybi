package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ybi.com/appbendaharaybi/auth"
	"ybi.com/appbendaharaybi/cashflow"
	"ybi.com/appbendaharaybi/handler"
	"ybi.com/appbendaharaybi/helper"
	"ybi.com/appbendaharaybi/user"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	
	userRepository:= user.NewRepository(db)
	cashflowRepository:= cashflow.NewRepository(db)


	userService := user.NewService(userRepository)
	cashflowService := cashflow.NewService(cashflowRepository)
	authService := auth.NewService()


	userHandler := handler.NewUserHandler(userService, authService)
	cashflowHandler:= handler.NewCashflowHandler(cashflowService)


	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	api.POST("/login", userHandler.Login)



	api.GET("/cashflows", authMiddleware(authService,userService), cashflowHandler.GetCashflow)
	api.POST("/createcashflow", authMiddleware(authService,userService), cashflowHandler.CreateCashflow)	
	api.PUT("/cashflow", authMiddleware(authService,userService), cashflowHandler.UpdateCashflow)
	api.DELETE("/cashflow", authMiddleware(authService,userService), cashflowHandler.DeleteCashflow)

	router.Run(":8080")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		// fmt.Println(claim)
		if !ok || !token.Valid {
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		user_ID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(user_ID)
		if err != nil {
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}
