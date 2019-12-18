package redis

import (
	"fmt"
	"time"

	"github.com/ckeyer/diego/pkgs/storage/metadata"
	"github.com/ckeyer/diego/types"
)

var _ metadata.ProjectStorager = &RedisStorage{}

// CreateProject
func (r *RedisStorage) GetProject(ns, name string) (*types.Project, error) {
	prj := &types.Project{}
	err := r.getKV(name, prj)
	if err != nil {
		return nil, err
	}
	return prj, nil
}

// CreateProject
func (r *RedisStorage) CreateProject(prj *types.Project) error {
	prj.Created = time.Now()
	if exi, err := r.ExistsNamespace(prj.Namespace); err != nil {
		return err
	} else if !exi {
		return metadata.ErrNotExists
	}

	if err := r.setKV(prjKey(prj.Namespace, prj.Name), prj, false); err != nil {
		return fmt.Errorf("create project %s/%s failed, %s", prj.Namespace, prj.Name, err)
	}
	return nil
}

// CreateProject
func (r *RedisStorage) ListProjects(ns string) ([]*types.Project, error) {
	ps := []*types.Project{}
	err := r.listKVs(prjKey(ns, "*"), &ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

// CreateProject
func (r *RedisStorage) RemoveProject(ns, name string) error {
	if err := r.deleteKey(prjKey(ns, name)); err != nil {
		return err
	}

	return nil
}
