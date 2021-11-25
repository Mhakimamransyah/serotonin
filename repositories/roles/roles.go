package roles

import (
	"time"

	"gorm.io/gorm"
)

type RolesTables struct {
	gorm.Model
	ID        int       `gorm:"id;primaryKey:autoIncrement"`
	Name      string    `gorm:"name;not null;type:varchar(100);uniqueIndex:Name"`
	CreatedAt time.Time `gorm:"created_at;type:datetime;default:null"`
	UpdatedAt time.Time `gorm:"updated_at;type:datetime;default:null"`
	DeletedAt time.Time `gorm:"deleted_at;type:datetime;default:null"`
}
