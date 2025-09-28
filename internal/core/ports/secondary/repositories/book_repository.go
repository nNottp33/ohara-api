package repositories

import (
	"github.com/nNottp33/ohara-api/internal/core/domain"
)

type BooksRepository interface {
	Save(books domain.Books) error
}
