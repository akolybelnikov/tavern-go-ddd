// Package entity holds all the entities that are shared across all the subdomains

package entity

import "github.com/google/uuid"

// Item represents an Item for all Subdomains
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
