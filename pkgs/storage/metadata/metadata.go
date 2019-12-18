package metadata

import (
	"errors"

	"github.com/ckeyer/diego/pkgs/users"
	"github.com/ckeyer/diego/types"
)

var (
	ErrNotExists = errors.New("not exists.")
)

type MetadataStorager interface {
	UserStorager
	OrgStorager
	NamespaceStorager
	ProjectStorager
	FileIndexer
}

type UserStorager interface {
	ListUsers(users.ListUserOption) ([]*users.User, error)
	GetUser(name string) (*users.User, error)
	CreateUser(*users.User) error
	UpdateUser(*users.User) (*users.User, error)
	RemoveUser(name string) error
}

type OrgStorager interface {
	ListOrgs(users.ListOrgOption) ([]*users.Org, error)
	GetOrg(name string) (*users.Org, error)
	CreateOrg(*users.Org) error
	UpdateOrg(*users.Org) (*users.Org, error)
	RemoveOrg(name string) error
}

// 创建 用户 和 组织 的时候，需要同时创建命名空间
type NamespaceStorager interface {
	ExistsNamespace(name string) (bool, error)
	GetNamespace(name string) (*users.Namespace, error)
	CreateNamespace(*users.Namespace) error
	UpdateNamespace(*users.Namespace) (*users.Namespace, error)
	RemoveNamespace(name string) error
}

type ProjectStorager interface {
	GetProject(namespace, name string) (*types.Project, error)
	ListProjects(namespace string) ([]*types.Project, error)
	CreateProject(*types.Project) error
	RemoveProject(namespace, name string) error
}

type FileIndexer interface {
}
