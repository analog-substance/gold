package types

import (
	"github.com/google/uuid"
	"sort"
	"strings"
)

type Human struct {
	UUID      string
	Names     []string
	Emails    []string
	Passwords []string
	Roles     []string
}

func NewHuman() *Human {
	return &Human{UUID: uuid.New().String()}
}

func (h *Human) AddEmail(email string) {
	email = strings.ToLower(strings.Trim(email, "“” "))

	for _, e := range h.Emails {
		if email == e {
			// we already have that email
			return
		}
	}
	h.Emails = append(h.Emails, strings.ToLower(email))
	sort.Strings(h.Emails)
}

func (h *Human) AddName(name string) {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return
	}

	for _, n := range h.Names {
		if strings.EqualFold(name, n) {
			// we already have that name
			return
		}
	}
	h.Names = append(h.Names, name)
	sort.Strings(h.Names)
}

func (h *Human) AddRole(roleName string) {
	roleName = strings.TrimSpace(roleName)
	if len(roleName) == 0 {
		return
	}

	for _, n := range h.Roles {
		if strings.EqualFold(roleName, n) {
			// we already have that name
			return
		}
	}
	h.Roles = append(h.Roles, roleName)
	sort.Strings(h.Roles)
}

func (h *Human) AddPassword(password string) {
	if len(password) == 0 {
		return
	}

	for _, p := range h.Passwords {
		if password == p {
			// we already have that password
			return
		}
	}
	h.Passwords = append(h.Passwords, password)
	sort.Strings(h.Passwords)
}

func (h *Human) Merge(otherHuman *Human) {
	for _, e := range otherHuman.Emails {
		h.AddEmail(e)
	}

	for _, n := range otherHuman.Names {
		h.AddName(n)
	}
}
