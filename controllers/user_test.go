package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/nuts300/test-echo/db"
	"github.com/nuts300/test-echo/models"
	"github.com/nuts300/test-echo/resources"
)

var database = db.GetDB()

var userJSON = `{"email":"user_test@test.com", "password": "1234"}`

type UserRes struct {
	Email string `json:"email"`
}

func generateContextAndResponse(path string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, path, strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	context := e.NewContext(req, recorder)
	return context, recorder
}

func TestCreateUser(t *testing.T) {
	c, rec := generateContextAndResponse("/users")

	userResource := resources.NewUserResource(database)
	userController := NewUserController(userResource)

	postUser := models.NewUser()
	if err := json.Unmarshal([]byte(userJSON), postUser); err != nil {
		assert.Fail(t, "Failed unmarshal post data.", err.Error())
	}

	if assert.NoError(t, userController.CreateUser(c)) {
		resUser := models.NewUser()
		if err := json.Unmarshal(rec.Body.Bytes(), resUser); err != nil {
			assert.Fail(t, "Failed unmarshal post data.", err.Error())
		} else {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, postUser.Email, resUser.Email)
		}
	}
}
