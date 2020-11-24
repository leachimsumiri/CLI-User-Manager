package model

import (
	"gorm.io/gorm"
)

const MIN_PASSWORD_LENGTH = 4

// Anmerkung (Stefan): Wenn Variablennamen hier klein geschrieben sind, sind sie "private" und können nicht mehr gesetzt werden, daher in "gorm" groß üblich
type User struct {
	gorm.Model // beinhaltet schon ID/primKey, CreatedAt, UpdatedAt, DeletedAt
	//ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	FirstName string // Nullable -> sql.NullString hier bei gorm, bzw weiter unten zB sql.NullTime relevant
	LastName  string
	Password  string
}

// --------------------------------------------------------------------------------------------------------------------

// // Voll! Folgendes wär auch möglich, aber etwas viel Boilerplate -> daher imho direkt in Controller (oder ggf UserController)
// // --------------------------------------------------------------------------------------------------------------------------
// // func (user *User) Create() *gorm.DB {
// // 	return Inst.Create(user)
// // }
