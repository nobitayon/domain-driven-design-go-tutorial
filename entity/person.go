// package entities holds all the entities that
// are shared accross subdomain

package entity

import "github.com/google/uuid"

// Person is entity that represents a person in
// all domain
type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
