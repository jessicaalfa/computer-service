package main

import (
	"tugas/configs"
	"tugas/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	configs.ConnectDB()
	// Route / to handler function
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBooksByIDController)
	e.POST("/books", controllers.CreateBooksController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
	
}
