package users

import (
	"time"

	"github.com/ckeyer/diego/pkgs/apis"
)

type Org struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Desc  string `json:"desc"`
	Owner string `json:"owner"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type Namespace struct {
	Name      string `json:"name"`
	OwnerType string `json:"owner_type"`
}

type ListOrgOption struct {
	apis.ListOption
}
