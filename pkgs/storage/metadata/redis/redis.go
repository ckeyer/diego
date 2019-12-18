package redis

import (
	"strings"

	"github.com/ckeyer/diego/pkgs/storage/metadata"
	"github.com/gomodule/redigo/redis"
)

type RedisStorage struct {
	redis.Conn
}

var _ metadata.MetadataStorager = &RedisStorage{}

func NewRedisStorager(conn redis.Conn) metadata.MetadataStorager {
	return &RedisStorage{
		Conn: conn,
	}
}

const (
	keySeparator = ":"

	keyPrefixNamespace = "ns"
	keyPrefixUser      = "user"
	keyPrefixOrg       = "org"
	keyPrefixProject   = "prj"
)

var (
	nsKey   = redisKeyFunc(keyPrefixNamespace)
	userKey = redisKeyFunc(keyPrefixUser)
	orgKey  = redisKeyFunc(keyPrefixOrg)
	prjKey  = func(ns, name string) string {
		return strings.Join([]string{keyPrefixProject, ns, name}, keySeparator)
	}
)

func redisKeyFunc(prefix string) func(string) string {
	return func(name string) string {
		return prefix + keySeparator + name
	}
}
