package expense

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type handler struct {
	DB *sql.DB
}

func NewApplication(db *sql.DB) *handler {
	return &handler{db}
}

func (h *handler) GetAllExpenses(c echo.Context) error {
	rows, err := h.DB.Query("SELECT * FROM expenses")
	if err != nil {
		return err
	}
	defer rows.Close()

	var nn = []Expense{}
	var n = Expense{}

	for rows.Next() {
		err := rows.Scan(&n.ID, &n.Title, &n.Amount, &n.Note, pq.Array(&n.Tags))
		if err != nil {
			log.Fatal(err)
		}
		nn = append(nn, n)
	}

	return c.JSON(http.StatusOK, nn)
}

func (h *handler) GetExpense(c echo.Context) error {
	expense := Expense{}
	var expenses = []Expense{}
	id := c.Param("id")
	stmt, err := h.DB.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")
	if err != nil {
		return err
	}

	row := stmt.QueryRow(id)
	err = row.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		log.Fatal(err)
	}
	expenses = append(expenses, expense)

	return c.JSON(http.StatusOK, expenses)
}

func GetExpenseHandler(c echo.Context) error {
	expense := Expense{}
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query expense detail:" + err.Error()})
	}

	row := stmt.QueryRow(id)
	err = row.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "expense not found"})
	case nil:
		return c.JSON(http.StatusOK, expense)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan expense:" + err.Error()})
	}
}

func GetAllExpensesHandler(c echo.Context) error {
	auth := c.Request().Header.Get("Authorization")
	if auth != "November 10, 2009wrong_token" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all expenses detail:" + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all expenses:" + err.Error()})
	}

	expenses := []Expense{}
	for rows.Next() {
		var ex Expense
		err = rows.Scan(&ex.ID, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan expense:" + err.Error()})
		}
		expenses = append(expenses, ex)
	}

	return c.JSON(http.StatusOK, expenses)
}

func GetExpense(db *sql.DB, id int) (Expense, error) {
	expense := Expense{}
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")
	if err != nil {
		log.Fatal("can't prepare query expense", err)
		return expense, err
	}

	row := stmt.QueryRow(id)
	err = row.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		log.Fatal("can't scan expense", err)
		return expense, err
	}

	return expense, nil
}

func GetAllExpenses(db *sql.DB) ([]Expense, error) {
	expenses := []Expense{}
	rows, err := db.Query("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		log.Fatal("can't query all expenses:", err)
		return expenses, err
	}

	for rows.Next() {
		ex := Expense{}
		err := rows.Scan(&ex.ID, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
		if err != nil {
			return expenses, err
		}
		expenses = append(expenses, ex)
	}

	return expenses, nil
}
