package main

import (
	"os"

	"github.com/invertedbit/gocook/database"
	"github.com/invertedbit/gocook/handlers"
	"github.com/invertedbit/gocook/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DSN string
}

func main() {

	godotenv.Load(".env.local")
	godotenv.Load()

	cfg := Config{
		DSN: os.Getenv("DSN"),
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Media{}, &models.Recipe{}, &models.InstructionStep{}, &models.Ingredient{}, &models.RecipeIngredient{})

	database.DBConn = db

	app := handlers.New()

	defer app.Shutdown()

	app.Listen(":3000")
}
