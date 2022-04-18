// Package entity holds all the entities that are shared across all Subdomains
package entity

import "github.com/google/uuid"

// Person is an entity that represents a person in all Domains
type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
