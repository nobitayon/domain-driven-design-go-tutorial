// package aggreage holds our aggregates that combine
// many entity into full object
package aggregate

import (
	"ddd/entity"
	"ddd/valueobject"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer need to have valid name")
)

type Customer struct {
	// person is the root entity of customer
	// which means person.ID is the main identifier for the customer
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new customer
// aggregate. It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetId() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {

	return c.person.Name
}
