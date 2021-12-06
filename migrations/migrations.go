package migrations

import (
	"serotonin/repositories/users"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&users.UsersTable{},
	)
}
