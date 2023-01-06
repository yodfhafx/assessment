package expense

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func UpdateExpenseHandler(c echo.Context) error {
	var ex Expense
	err := c.Bind(&ex)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	stmt, err := db.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare expense update:" + err.Error()})
	}

	if _, err := stmt.Exec(&ex.ID, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags)); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "error execute update:" + err.Error()})
	}

	return c.JSON(http.StatusCreated, ex)
}

func UpdateExpense(db *sql.DB, ex Expense) (Expense, error) {
	expense := Expense{}
	row := db.QueryRow(`UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id = $1  RETURNING id,title,amount,note,tags`, ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags))

	err := row.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		log.Fatal("can't scan expense", err)
		return expense, err
	}

	return expense, nil
}
