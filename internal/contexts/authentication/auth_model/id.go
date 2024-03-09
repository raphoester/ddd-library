package auth_model

import (
	"errors"

	"github.com/google/uuid"
)

type ID uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}

func (id ID) Validate() error {
	if uuid.UUID(id) == uuid.Nil {
		return errors.New("id is empty")
	}
	return nil
}
