package users

import (
	"time"

	"github.com/ckeyer/diego/pkgs/apis"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name  string `json:"name"`
	Email string `json:"email"`
	Desc  string `json:"desc"`

	Joined  time.Time `json:"joined"`
	Updated time.Time `json:"updated"`
}

type ListUserOption struct {
	apis.ListOption
}
