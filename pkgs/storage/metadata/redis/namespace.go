package redis

import (
	"fmt"

	"github.com/ckeyer/diego/pkgs/storage/metadata"
	"github.com/ckeyer/diego/types"
)

var _ metadata.NamespaceStorager = &RedisStorage{}

func (r *RedisStorage) ExistsNamespace(name string) (bool, error) {
	return r.existsKey(nsKey(name))
}

func (r *RedisStorage) GetNamespace(name string) (*types.Namespace, error) {
	ns := &types.Namespace{}
	if err := r.getKV(nsKey(name), &ns); err != nil {
		return nil, fmt.Errorf("get namespace %s failed, %s", name, err)
	}

	return ns, nil
}

func (r *RedisStorage) CreateNamespace(ns *types.Namespace) error {
	if err := r.setKV(nsKey(ns.Name), ns, false); err != nil {
		return fmt.Errorf("create namespace %s failed, %s", ns.Name, err)
	}
	return nil
}

func (r *RedisStorage) UpdateNamespace(ns *types.Namespace) (*types.Namespace, error) {
	if err := r.setKV(nsKey(ns.Name), ns, true); err != nil {
		return nil, fmt.Errorf("update namespace %s failed, %s", ns.Name, err)
	}
	return ns, nil
}

func (r *RedisStorage) RemoveNamespace(name string) error {
	if err := r.deleteKey(nsKey(name)); err != nil {
		return fmt.Errorf("remove namespace %s failed, %s", name, err)
	}
	return nil
}
