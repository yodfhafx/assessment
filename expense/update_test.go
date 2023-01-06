//go:build unit
// +build unit

package expense

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestUpdateExpense(t *testing.T) {
	expense := Expense{Title: "Golang", Amount: 300.00, Note: "simple", Tags: []string{"banana"}}
	db, mock, _ := sqlmock.New()
	rows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).
		AddRow(1, "Golang", 300.00, "simple", pq.Array([]string{"banana"}))

	mock.ExpectQuery("UPDATE expenses").
		WithArgs("Golang", 300.00, "simple", pq.Array([]string{"banana"})).
		WillReturnRows(rows)

	result, err := UpdateExpense(db, expense)

	if assert.NoError(t, err) {
		assert.EqualValues(t, result.Title, expense.Title)
		assert.EqualValues(t, result.Amount, expense.Amount)
		assert.EqualValues(t, result.Note, expense.Note)
		assert.EqualValues(t, result.Tags, expense.Tags)
	}
}
