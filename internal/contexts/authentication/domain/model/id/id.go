package id

import (
	"errors"

	"github.com/google/uuid"
)

type ID uuid.UUID

func Create() ID {
	return ID(uuid.New())
}

func NewFromString(str string) (*ID, error) {
	uid, err := uuid.Parse(str)
	if err != nil {
		return nil, err
	}

	return (*ID)(&uid), nil
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
