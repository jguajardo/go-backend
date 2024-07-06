package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	// Intentar cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables instead")
	}

	// Obtener la variable de entorno GO_ENV
	env := os.Getenv("GO_ENV")
	if env == "" {
		log.Fatal("GO_ENV is not set")
	}

	// Verificar si estamos en un entorno local o en un entorno de producción
	if env == "local" {
		log.Println("Running in local environment")
	} else if env == "production" {
		log.Println("Running in production environment")
	}

	// Obtener las variables de entorno para la base de datos
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	// Verificar que todas las variables de entorno necesarias están configuradas
	if user == "" || password == "" || host == "" || port == "" || dbName == "" || sslmode == "" || timezone == "" {
		log.Fatalf("One or more environment variables are missing. Ensure all required variables are set.")
	}

	// Construir el DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbName, port, sslmode, timezone)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}
