package borrow_model

import "github.com/google/uuid"

type Borrower struct {
	id     uuid.UUID
	userID string
}
