package users

import (
	"errors"
	"slices"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

var validRoles = []Role{RoleAdmin, RoleUser}

func (r Role) Validate() error {
	if !slices.Contains(validRoles, r) {
		return errors.New("invalid role")
	}

	return nil
}
