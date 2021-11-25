package migrations

import (
	"serotonin/repositories/address"
	"serotonin/repositories/roles"
	"serotonin/repositories/users"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&roles.RolesTables{},
		&users.UsersTable{},
		&address.AddressTable{},
	)
}
