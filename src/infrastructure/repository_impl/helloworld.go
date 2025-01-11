package repository_impl

import (
	"context"
	"database/sql"
	"errors"
	"go-ddd-template/src/domain/entity"
	"go-ddd-template/src/domain/repository"
	"go-ddd-template/src/infrastructure"
	"log"
	"time"
)

type HelloRepository struct {
	infrastructure.SqlHandler
}

func NewHelloRepository(sqlHandler infrastructure.SqlHandler) repository.HelloRepository {
	helloRepository := HelloRepository{sqlHandler}
	return &helloRepository
}

func (helloRepo *HelloRepository) Find(ctx context.Context,id int) (hello entity.Hello, err error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	conn := helloRepo.SqlHandler.Conn
	if conn == nil {
		log.Printf("[Error]HelloRepository: Database connection is nil")
		err = repository.ErrDatabaseUnavailable
		return
	}
	row := conn.QueryRowContext(ctx, "SELECT id, name, tag FROM hello_world WHERE id = ?", id)
	err = row.Scan(&hello.Id, &hello.Name, &hello.Tag)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("[Error]HelloRepository: Database connection timeout")
			err = repository.ErrDatabaseUnavailable
			return
		}
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
