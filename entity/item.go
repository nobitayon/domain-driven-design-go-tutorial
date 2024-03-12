// package entities holds all the entities that
// are shared accross subdomain

package entity

import "github.com/google/uuid"

// Item is entity that represents a item in
// all domain
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
