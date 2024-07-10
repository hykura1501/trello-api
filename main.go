package main

import (
	"log"
	"net/http"
	"os"
	"trello-api/database"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Connect to Database
	sql := &database.SQL{
		DataSourceName: os.Getenv("DB_CONNECTION_STRING"),
	}
	sql.Connect()
	defer sql.Close()

	//Echo init
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	PORT := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + PORT))
}
