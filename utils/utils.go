package utils

import "github.com/ziauddinmohammad/basic-go-CRUD-service/models"

func NextID() int64 {
	index := len(models.Books)
	if index == 0 {
		return 1
	}
	nextid := models.Books[index-1].ID + 1
	return nextid
}

func IDExists(id int64) int {
	for index, book := range models.Books {
		if book.ID == id {
			return index
		}
	}
	return -1
}
