package borrow_model

import (
	"time"

	"github.com/google/uuid"
)

type Suggestion struct {
	id         uuid.UUID
	authoredBy uuid.UUID // borrower id
	bookTitle  string
	comment    string
	createdAt  time.Time
}
