package address

import (
	"serotonin/repositories/users"

	"gorm.io/gorm"
)

type AddressTable struct {
	gorm.Model
	ID          int              `gorm:"id;primaryKey:autoIncrement"`
	Street      string           `gorm:"street;not null;type:varchar(100)"`
	City        string           `gorm:"city;not null;type:varchar(100)"`
	Province    string           `gorm:"province;not null;type:varchar(100)"`
	Postal_Code string           `gorm:"postal_code;not null;type:varchar(100)"`
	UsersId     int              `gorm:"users_id;not null;type:bigint"`
	User        users.UsersTable `gorm:"foreignKey:UsersId"`
}
