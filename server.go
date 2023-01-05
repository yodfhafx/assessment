package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yodfhafx/assessment/expense"
)

func main() {
	expense.InitDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/expenses", expense.CreateExpensesHandler)
	e.GET("/expenses/:id", expense.GetExpenseHandler)
	log.Fatal(e.Start(os.Getenv("PORT")))
}
