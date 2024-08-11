package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"rental-mobil/controllers"
	"rental-mobil/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", dbConfig)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	database.DBMigrate(DB)

	router := gin.Default()

	// Rute publik
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Middleware otentikasi

	// Rute yang memerlukan autentikasi
	protected := router.Group("/api", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	{
		protected.GET("/cars", controllers.GetAllCars)
		protected.POST("/cars", controllers.InsertCar)
		protected.PUT("/cars/:id", controllers.UpdateCar)
		protected.DELETE("/cars/:id", controllers.DeleteCar)

		protected.GET("/rentals", controllers.GetAllRentals)
		protected.POST("/rentals", controllers.InsertRental)
		protected.PUT("/rentals/:id", controllers.UpdateRental)
		protected.DELETE("/rentals/:id", controllers.DeleteRental)

		protected.GET("/payments", controllers.GetAllPayments)
		protected.POST("/payments", controllers.InsertPayment)
		protected.PUT("/payments/:id", controllers.UpdatePayment)
		protected.DELETE("/payments/:id", controllers.DeletePayment)
	}

	router.Run(":8080")
}

// AuthMiddleware is used to check the token from Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Token biasanya diawali dengan "Bearer ", kita ambil bagian setelah spasi
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parsing token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Pastikan metode signing adalah HS256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNoLocation
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Token valid, lanjutkan permintaan
		c.Next()
	}
}
