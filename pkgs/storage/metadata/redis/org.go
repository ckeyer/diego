package redis

import (
	"fmt"
	"time"

	"github.com/ckeyer/diego/pkgs/storage/metadata"
	"github.com/ckeyer/diego/types"
)

var _ metadata.OrgStorager = &RedisStorage{}

// GetOrg return user by name
func (r *RedisStorage) GetOrg(name string) (*types.Org, error) {
	org := &types.Org{}
	if err := r.getKV(orgKey(name), &org); err != nil {
		return nil, fmt.Errorf("get org %s failed, %s", name, err)
	}

	return org, nil
}

func (r *RedisStorage) CreateOrg(o *types.Org) error {
	o.Created = time.Now()
	o.Updated = time.Now()

	ns := &types.Namespace{
		Name:      o.Name,
		OwnerType: types.OwnerTypeOrg,
	}
	if err := r.CreateNamespace(ns); err != nil {
		return err
	}

	if err := r.setKV(orgKey(o.Name), o, false); err != nil {
		// rollback
		r.RemoveNamespace(ns.Name)
		return fmt.Errorf("create org %s failed, %s", o.Name, err)
	}

	return nil
}

func (r *RedisStorage) UpdateOrg(org *types.Org) (*types.Org, error) {
	org.Updated = time.Now()
	if err := r.setKV(orgKey(org.Name), org, true); err != nil {
		return nil, fmt.Errorf("update org %s failed, %s", org.Name, err)
	}
	return org, nil
}

// ListOrgs
func (r *RedisStorage) ListOrgs(opt types.ListOrgOption) ([]*types.Org, error) {
	orgs := []*types.Org{}
	err := r.listKVs(orgKey("*"), &orgs)
	if err != nil {
		return nil, err
	}

	return orgs, nil
}

// RemoveOrg
func (r *RedisStorage) RemoveOrg(name string) error {
	// TODO: move all projects in orgs.

	if err := r.deleteKey(orgKey(name)); err != nil {
		return err
	}

	if err := r.RemoveNamespace(name); err != nil {
		return err
	}

	return nil
}
