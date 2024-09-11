package models

import (
	"gorm.io/gorm"
)

// Enum types

type Role int
type Status int
type UserAccountType int

const (
	Admin      Role = 0
	TeamMember Role = 1

	Pending  Status = 0
	Inactive Status = 1
	Active   Status = 2

	Individual UserAccountType = 0
	Enterprise UserAccountType = 1
)

type User struct {
	gorm.Model
	Name            string `gorm:"size:155" json:"name"`
	Email           string `gorm:"size:155;unique" json:"email"`
	Role            Role
	Status          Status
	UserAccountType UserAccountType
	DeletedAt       gorm.DeletedAt `gorm:"index"` // Soft delete field
}
