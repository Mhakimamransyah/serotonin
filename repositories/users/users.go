package users

import (
	"serotonin/business/users"
	"serotonin/repositories/roles"
	"time"

	"gorm.io/gorm"
)

type UsersRepository struct {
	DB *gorm.DB
}

func InitRepository(DB *gorm.DB) *UsersRepository {
	return &UsersRepository{
		DB: DB,
	}
}

type UsersTable struct {
	gorm.Model
	ID        int               `gorm:"id;primaryKey:autoIncrement"`
	Name      string            `gorm:"name;not null;type:varchar(100);uniqueIndex:Name"`
	Email     string            `gorm:"email;not null;type:varchar(100);uniqueIndex:Email"`
	Username  string            `gorm:"username;not null;type:varchar(100);uniqueIndex:Username"`
	Password  string            `gorm:"password;not null;type:text"`
	Phone     string            `gorm:"phone;not null;type:varchar(100);uniqueIndex:Phone"`
	RolesId   int               `gorm:"roles_id;not null;type:bigint;"`
	CreatedAt time.Time         `gorm:"created_at;type:datetime;default:null"`
	UpdatedAt time.Time         `gorm:"updated_at;type:datetime;default:null"`
	DeletedAt time.Time         `gorm:"deleted_at;type:datetime;default:null"`
	Roles     roles.RolesTables `gorm:"foreignKey:RolesId"`
}

func (repo *UsersRepository) CreateUser(user *users.Users) error {
	return nil
}
func (repo *UsersRepository) Login(username, password string) (*users.Users, error) {
	return nil, nil
}
func (repo *UsersRepository) Get(username string) (*users.Users, error) {
	return nil, nil
}
func (repo *UsersRepository) Update(user *users.Users) error {
	return nil
}
func (repo *UsersRepository) Delete(user *users.Users) error {
	return nil
}
