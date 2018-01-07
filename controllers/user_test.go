package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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

var createdUser models.User

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

func generateUserController() UserController {
	userResource := resources.NewUserResource(database)
	userController := NewUserController(userResource)
	return userController
}

func TestCreateUser(t *testing.T) {
	c, rec := generateContextAndResponse(echo.POST, "/users", &userJSON)
	userController := generateUserController()

	postUser := models.NewUser()
	if err := json.Unmarshal([]byte(userJSON), postUser); err != nil {
		assert.Fail(t, "Failed unmarshal post data.", err.Error())
	}

	if assert.NoError(t, userController.CreateUser(c)) {
		resUser := models.NewUser()
		if err := json.Unmarshal(rec.Body.Bytes(), resUser); err != nil {
			assert.Fail(t, "Failed unmarshal post data.", err.Error())
		} else {
			createdUser = *resUser
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, postUser.Email, resUser.Email)
		}
	}
}

func TestGetUser(t *testing.T) {
	c, rec := generateContextAndResponse(echo.GET, "/", nil)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(createdUser.ID))
	userController := generateUserController()

	if assert.NoError(t, userController.GetUser(c)) {
		resUser := models.NewUser()
		if err := json.Unmarshal(rec.Body.Bytes(), resUser); err != nil {
			assert.Fail(t, "Failed unmarshal post data.", err.Error())
		} else {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, createdUser.Email, resUser.Email)
			assert.Equal(t, createdUser.ID, resUser.ID)
		}
	}
}
