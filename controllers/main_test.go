package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo"
)

func generateContextAndResponse(method string, path string, payLoad *string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	var req *http.Request
	if payLoad != nil {
		req = httptest.NewRequest(method, path, strings.NewReader(*payLoad))
	} else {
		req = httptest.NewRequest(method, path, nil)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	context := e.NewContext(req, recorder)
	return context, recorder
}
