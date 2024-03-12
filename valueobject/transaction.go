package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transactions is a valueobject because it has
// no identifier and immutable
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
