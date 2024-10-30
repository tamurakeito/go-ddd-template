package infrastructure

import (
	"log"

	"go-ddd-template/src/domain/model"
	"go-ddd-template/src/domain/repository"
)

type HelloRepository struct {
	SqlHandler
}

func NewHelloRepository(sqlHandler SqlHandler) repository.HelloRepository {
	helloRepository := HelloRepository{sqlHandler}
	return &helloRepository
}

func (helloRepo *HelloRepository) Find(id int) (hello model.Hello, err error) {
	row := helloRepo.SqlHandler.Conn.QueryRow("SELECT id, name, tag FROM hello_world WHERE id = ?", id)
	err = row.Scan(&hello.ID, &hello.Name, &hello.Tag)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
