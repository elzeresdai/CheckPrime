package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCheckPrimesPass(t *testing.T) {
	str := `{"numbers":[2,44,7,1]}`
	expected := `["true","false","true","false"]` + "\n"

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/check", strings.NewReader(str))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := CheckPrimes(c)

	if assert.NoError(t, h) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestCheckPrimesFail(t *testing.T) {
	str := `{"numbers":[2,"error",7,1]}`
	expected := `"the given message is invalid in index: 1"` + "\n"

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/check", strings.NewReader(str))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := CheckPrimes(c)

	if assert.NoError(t, h) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
