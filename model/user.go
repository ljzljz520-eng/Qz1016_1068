package model

import "errors"

var ErrInvalidStatus = errors.New("invalid status")
var ErrNotFound = errors.New("not found")
var ErrConflict = errors.New("version conflict")

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Active bool   `json:"active"`
}

func (u User) CanReview() bool { return u.Active && (u.Role == "manager" || u.Role == "auditor") }
func (u User) CanAssign() bool { return u.Active && u.Role != "viewer" }
