package authHandler

import (
	"bytes"
	"encoding/json"
	"example/gin-webserver/database"
	"example/gin-webserver/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	database.SetupMockDB()
	defer database.TeardownMockDB()

	r := gin.Default()
	r.POST("/register", Register)

	t.Run("Valid Registration", func(t *testing.T) {
		validInput := model.UserInput{
			Username: "testuser",
			Password: "testpw",
		}
		inputJSON, _ := json.Marshal(validInput)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(inputJSON))

		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		// You can further check the response body or other aspects of the response as needed.
		// For example, you can unmarshal the response body to verify its contents.
		responseBody := w.Body.String()
		var responseJSON map[string]interface{}
		json.Unmarshal([]byte(responseBody), &responseJSON)
		userMap, _ := responseJSON["user"].(map[string]interface{})
		assert.Equal(t, "testuser", userMap["username"])
	})

	t.Run("Invalid Registration", func(t *testing.T) {
		validInput := model.UserInput{
			Password: "testpw",
		}
		inputJSON, _ := json.Marshal(validInput)

		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(inputJSON))

		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
