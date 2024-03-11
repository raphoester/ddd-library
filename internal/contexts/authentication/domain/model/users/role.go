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

func (r Role) validate() error {
	if !slices.Contains(validRoles, r) {
		return errors.New("invalid role")
	}

	return nil
}

func (r Role) String() string {
	return string(r)
}
