package storage

import "book-catalog/models"

var Books = make(map[string]models.Book)
var LastID int = 0