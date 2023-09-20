package entryController

import (
	"example/gin-webserver/database"
	"example/gin-webserver/model"
)

func Create(entry model.Entry) (model.Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return entry, err
	}
	return entry, nil
}
