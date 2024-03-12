package product

import (
	"ddd/aggregate"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("no such product")
	ErrProductAlreadyExists = errors.New("there is already such product")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
