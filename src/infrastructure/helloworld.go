package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
	"log"
)

type HelloRepository struct {
	SqlHandler
}

func NewHelloRepository(sqlHandler SqlHandler) repository.HelloRepository {
	helloRepository := HelloRepository{sqlHandler}
	return &helloRepository
}

func (helloRepo *HelloRepository) checkConnection() error {
	if helloRepo.SqlHandler.Conn == nil {
		log.Printf("[Error]HelloRepository: Database connection is nil")
		return repository.ErrDatabaseUnavailable
	}
	return nil
}

func (helloRepo *HelloRepository) Find(ctx context.Context,id int) (hello entity.Hello, err error) {
	if err = helloRepo.checkConnection(); err != nil {
		return
	}
	row := helloRepo.SqlHandler.Conn.QueryRow("SELECT id, name, tag FROM hello_world WHERE id = ?", id)
	err = row.Scan(&hello.Id, &hello.Name, &hello.Tag)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = repository.ErrResourceNotFound
			return
		}
		log.Printf("[Error]HelloRepository.Find: %v", err)
		err = repository.ErrInternal
		return
	}
	return
}
