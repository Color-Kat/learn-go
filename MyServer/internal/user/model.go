package user

import "gorm.io/gorm"

type User struct {
	Email    string `json:"email" gorm:"index"`
	Password string `json:"password"`
	Name     string `json:"name"`
	gorm.Model
}

func (u *User) NewUser() {

}
