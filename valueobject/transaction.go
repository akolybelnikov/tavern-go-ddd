// Package valueobject holds the Value Objects interfaces of the Domain
package valueobject

import (
	"github.com/google/uuid"
	"time"
)

// Transaction is a payment between two parties
type Transaction struct {
	// transaction values are immutable
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
