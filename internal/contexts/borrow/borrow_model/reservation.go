package borrow_model

import (
	"time"

	"github.com/google/uuid"
)

type reservation struct {
	id          uuid.UUID
	bookID      uuid.UUID
	userID      uuid.UUID
	createdAt   time.Time
	cancelledAt time.Time
}
