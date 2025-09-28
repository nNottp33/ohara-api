package services

import (
	"github.com/nNottp33/ohara-api/internal/core/domain"
	"github.com/nNottp33/ohara-api/internal/core/ports/secondary/repositories"
)

type BooksService interface {
	Create(books domain.Books) error
}

type BooksServiceImplement struct {
	repo repositories.BooksRepository
}
