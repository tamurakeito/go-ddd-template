package entity

import "go-ddd-template/src/domain/model"

type HelloWorld struct {
	ID    int           `json:"id"`
	Hello []model.Hello `json:"hello"`
}
