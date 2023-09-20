package entryController

import (
	"example/gin-webserver/database"
	"example/gin-webserver/model"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEntry(t *testing.T) {
	database.SetupMockDB()
	defer database.TeardownMockDB()

	// populate database
	testUser := model.User{Username: "test", Password: "test"}
	database.Database.Create(&testUser)
	log.Println("DB populated")

	// test case 1: create entry
	entry := model.Entry{Content: "test", UserID: 1}
	result, err := Create(entry)

	assert.NoError(t, err)
	assert.Equal(t, "test", result.Content)

	// test case 2: id does not exist
	entry = model.Entry{Content: "test2", UserID: 0}
	result, err = Create(entry)

	assert.Error(t, err)
}
