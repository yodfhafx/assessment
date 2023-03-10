package expense

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func CreateExpenseHandler(c echo.Context) error {
	var ex Expense
	err := c.Bind(&ex)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id", ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags))
	err = row.Scan(&ex.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, ex)
}

func CreateExpenses(db *sql.DB, ex Expense) (Expense, error) {
	row := db.QueryRow(`INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id,title,amount,note,tags`, ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags))
	expense := Expense{}

	err := row.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		log.Fatal("can't scan expense", err)
	}

	return expense, nil
}

func (h *handler) CreateExpense(c echo.Context) error {
	ex := Expense{}
	var expense = []Expense{}
	row := h.DB.QueryRow("INSERT INTO expenses (id, title, amount, note, tags) values ($1, $2, $3, $4, $5)  RETURNING id,title,amount,note,tags", 2, "Golang", 200, "simple", pq.Array([]string{"banana"}))

	err := row.Scan(&ex.ID, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
	if err != nil {
		log.Fatal(err)
	}
	expense = append(expense, ex)

	return c.JSON(http.StatusOK, expense)
}
