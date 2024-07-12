package main

import (
	"log"
	"net/http"
	"os"
	"trello-api/database"
	"trello-api/routes"

	"trello-api/handlers"
	repositoryImpl "trello-api/repository/repository_impl"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
		AllowCredentials: true,
	}))
	//Handler
	userHandler := handlers.UserHandler{
		UserRepo: repositoryImpl.NewUserRepository(sql),
	}
	boardHandler := handlers.BoardHandler{
		BoardRepo: repositoryImpl.NewBoardRepository(sql),
	}
	columnHandler := handlers.ColumnHandler{
		ColumnRepo: repositoryImpl.NewColumnRepository(sql),
	}
	//End Handler
	// Setup Router
	api := routes.API{
		Echo:          e,
		UserHandler:   userHandler,
		BoardHandler:  boardHandler,
		ColumnHandler: columnHandler,
	}
	api.SetupRouter()
	// Setup Router
	PORT := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + PORT))
}
