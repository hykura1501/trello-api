package main

import (
	"log"
	"os"
	"trello-api/database"
	"trello-api/routes"

	"trello-api/handlers"
	repositoryImpl "trello-api/repository/repository_impl"

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
	//Handler
	userHandler := handlers.UserHandler{
		UserRepo: repositoryImpl.NewUserRepository(sql),
	}
	//End Handler
	// Setup Router
	api := routes.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()
	// Setup Router
	PORT := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + PORT))
}
