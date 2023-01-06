//go:build integration
// +build integration

package expense

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestITCreateExpense(t *testing.T) {
	eh := echo.New()
	go func(e *echo.Echo) {
		db, err := sql.Open("postgres", "postgresql://root:root@db/go-example-db?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		h := NewApplication(db)

		e.POST("/expenses", h.CreateExpense)
		e.Start(fmt.Sprintf(":%d", 2565))
	}(eh)
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", 2565), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}

	reqBody := ``
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/expenses", 2565), strings.NewReader(reqBody))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := http.Client{}

	resp, err := client.Do(req)
	assert.NoError(t, err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println("resp: ", resp)
	assert.NoError(t, err)
	resp.Body.Close()

	expected := "[{\"id\":2,\"title\":\"Golang\",\"amount\":200,\"note\":\"simple\",\"tags\":[\"banana\"]}]"

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expected, strings.TrimSpace(string(byteBody)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}
