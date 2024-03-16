package controller

import (
	".com/internal/db"
)

type PersonController struct {
	personStore db.PersonStore
}

func NewPersonController(store db.PersonStore) *PersonController {
	return &PersonController{
		personStore: store,
	}
}
