package redis

import (
	"testing"
	"time"

	"github.com/ckeyer/diego/types"
	"github.com/gomodule/redigo/redis"
)

// TestRedisConn ...
func TestRedisConn(t *testing.T) {
	return
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		t.Error(err)
		return
	}

	mss, err := redis.Strings(conn.Do(rCmdMget, "user:ck2", "user:ck"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Error(mss)
	return

	_, err = conn.Do("SET", "hi", "hello2")
	if err != nil {
		t.Error(err)
		return
	}

	rcli := NewRedisStorager(conn)

	u := &types.User{
		Name:   "ck",
		Email:  "me2@ckeyer.com",
		Desc:   "hello",
		Joined: time.Now(),
	}

	if err := rcli.CreateUser(u); err != nil {
		t.Error(err)
		return
	}

	ur, err := rcli.GetUser("ck")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", ur)

	t.Error("...")
}
