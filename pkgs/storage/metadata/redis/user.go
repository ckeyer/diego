package redis

import (
	"fmt"
	"time"

	"github.com/ckeyer/diego/pkgs/storage/metadata"
	"github.com/ckeyer/diego/types"
)

var _ metadata.UserStorager = &RedisStorage{}

// GetUser return user by name
func (r *RedisStorage) GetUser(name string) (*types.User, error) {
	user := &types.User{}
	if err := r.getKV(userKey(name), &user); err != nil {
		return nil, fmt.Errorf("get user %s failed, %s", name, err)
	}

	return user, nil
}

func (r *RedisStorage) CreateUser(user *types.User) error {
	user.Joined = time.Now()
	user.Updated = time.Now()

	ns := &types.Namespace{
		Name:      user.Name,
		OwnerType: types.OwnerTypeUser,
	}
	if err := r.CreateNamespace(ns); err != nil {
		return err
	}

	if err := r.setKV(userKey(user.Name), user, false); err != nil {
		// rollback
		r.RemoveNamespace(ns.Name)
		return fmt.Errorf("create user %s failed, %s", user.Name, err)
	}

	return nil
}

func (r *RedisStorage) UpdateUser(user *types.User) (*types.User, error) {
	user.Updated = time.Now()
	if err := r.setKV(userKey(user.Name), user, true); err != nil {
		return nil, fmt.Errorf("update user %s failed, %s", user.Name, err)
	}
	return user, nil
}

// ListUsers
func (r *RedisStorage) ListUsers(opt types.ListUserOption) ([]*types.User, error) {
	users := []*types.User{}
	err := r.listKVs(userKey("*"), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// RemoveUser
func (r *RedisStorage) RemoveUser(name string) error {
	// TODO: move all projects in users.

	if err := r.deleteKey(userKey(name)); err != nil {
		return err
	}

	if err := r.RemoveNamespace(name); err != nil {
		return err
	}

	return nil
}
