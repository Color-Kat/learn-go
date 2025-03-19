package user

import "demo/http/pkg/database"

type UserRepository struct {
	database *database.Db
}

func NewUserRepository(db *database.Db) *UserRepository {
	return &UserRepository{database: db}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.database.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.database.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
