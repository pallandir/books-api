package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pallandir/books-api/router"
	"github.com/stretchr/testify/assert"
)

func TestBooksRoute(t *testing.T) {
	router := router.ApiRouter()

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	router.ServeHTTP(writer, req)

	assert.Equal(t, 200, writer.Code)
	assert.Contains(t, writer.Body.String(), "9781612680194")
	assert.Contains(t, writer.Body.String(), "9781781257654")
	assert.Contains(t, writer.Body.String(), "9780593419052")
}
