package authController

import (
	"example/gin-webserver/database"
	"example/gin-webserver/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	database.SetupMockDB()
	defer database.TeardownMockDB()

	// test case 1: create user
	user := model.User{Username: "test", Password: "testpw"}
	result, err := CreateUser(user)

	assert.NoError(t, err)
	assert.Equal(t, "test", result.Username)

	// test case 2: dupliate username
	user2 := model.User{Username: "test", Password: "testpw"}
	_, err = CreateUser(user2)

	assert.Error(t, err)
}

func TestFindUserById(t *testing.T) {
	database.SetupMockDB()
	defer database.TeardownMockDB()

	// populate database

	// test case 1: find user that exists

	// test case 2: find user that does not exist

}

func TestFindUserByUsername(t *testing.T) {
	database.SetupMockDB()
	defer database.TeardownMockDB()

	// populate database

	// test case 1: find user that exists

	// test case 2: find user that does not exist

}
