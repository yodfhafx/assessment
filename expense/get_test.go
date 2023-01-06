package expense

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetExpense(t *testing.T) {
	db, mock, _ := sqlmock.New()
	row := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(1, "Golang", 300.00, "simple", pq.Array([]string{"banana"}))

	mock.ExpectPrepare("SELECT id, title, amount, note, tags FROM expenses").
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(row)

	result, err := GetExpense(db, 1)
	if assert.NoError(t, err) {
		assert.EqualValues(t, result.Title, "Golang")
		assert.EqualValues(t, result.Amount, 300.00)
		assert.EqualValues(t, result.Note, "simple")
		assert.EqualValues(t, result.Tags, []string{"banana"})
	}
}

func TestGetAllExpenses(t *testing.T) {
	db, mock, _ := sqlmock.New()
	rows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(1, "Golang", 300.00, "simple", pq.Array([]string{"banana"})).
		AddRow(2, "Golang", 200.00, "relax", pq.Array([]string{"mango"})).
		AddRow(3, "Golang", 250.00, "nice", pq.Array([]string{"apple"}))

	mock.ExpectQuery("SELECT id, title, amount, note, tags FROM expenses").
		WillReturnRows(rows)

	result, err := GetAllExpenses(db)

	if assert.NoError(t, err) {
		assert.EqualValues(t, len(result), 3)
	}
}
