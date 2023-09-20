package entryHandler

import (
	"bytes"
	"encoding/json"
	"example/gin-webserver/database"
	"example/gin-webserver/model"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostEntry(t *testing.T) {
	database.SetupMockDB()
	defer database.TeardownMockDB()

	testUser := model.User{Username: "test", Password: "test"}
	database.Database.Create(&testUser)
	log.Println("db populated")

	r := gin.Default()
	r.POST("/entry", PostEntry)

	t.Run("Valid Entry", func(t *testing.T) {
		validInput := model.Entry{
			Content: "testContent",
			UserID:  1,
		}
		inputJSON, _ := json.Marshal(validInput)

		req, err := http.NewRequest("POST", "/entry", bytes.NewBuffer(inputJSON))

		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		// You can further check the response body or other aspects of the response as needed.
		// For example, you can unmarshal the response body to verify its contents.
		responseBody := w.Body.String()
		var responseJSON map[string]interface{}
		json.Unmarshal([]byte(responseBody), &responseJSON)
		entryMap, _ := responseJSON["data"].(map[string]interface{})
		assert.Equal(t, "testContent", entryMap["content"])
	})
}
