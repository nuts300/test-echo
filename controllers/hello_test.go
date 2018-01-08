package controllers

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func generateHelloController() HelloController {
	return NewHelloController()
}

func TestHello(t *testing.T) {
	c, rec := generateContextAndResponse(echo.GET, "/", nil)
	c.SetPath("/hello")
	helloController := generateHelloController()

	if assert.NoError(t, helloController.Hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
