package metadata

import (
	"errors"

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
	ListUsers(types.ListUserOption) ([]*types.User, error)
	GetUser(name string) (*types.User, error)
	CreateUser(*types.User) error
	UpdateUser(*types.User) (*types.User, error)
	RemoveUser(name string) error
}

type OrgStorager interface {
	ListOrgs(types.ListOrgOption) ([]*types.Org, error)
	GetOrg(name string) (*types.Org, error)
	CreateOrg(*types.Org) error
	UpdateOrg(*types.Org) (*types.Org, error)
	RemoveOrg(name string) error
}

// 创建 用户 和 组织 的时候，需要同时创建命名空间
type NamespaceStorager interface {
	ExistsNamespace(name string) (bool, error)
	GetNamespace(name string) (*types.Namespace, error)
	CreateNamespace(*types.Namespace) error
	UpdateNamespace(*types.Namespace) (*types.Namespace, error)
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
