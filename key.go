package gredis

import (
	"github.com/mediocregopher/radix/v3"
	gconv "github.com/og/x/conv"
	ge "github.com/og/x/error"
	"time"
)

func (self RedisClient) Exists(key string) (exists bool) {
	data := radix.MaybeNil{Rcv: &exists}
	err := self.Core.Do(radix.Cmd(&data, EXISTS, key)); ge.Check(err)
	return
}
func (self RedisClient) Expire(key string, second int) {
	err := self.Core.Do(radix.Cmd(nil, EXPIRE, key, gconv.IntString(second))); ge.Check(err)
}
func (self RedisClient) ExpireAt(key string, at time.Time) {
	err := self.Core.Do(radix.Cmd(nil, EXPIREAT, key, gconv.Int64String(at.Unix()))); ge.Check(err)
}

func (self RedisClient) Pexpire(key string, duration time.Duration) {
	err := self.Core.Do(radix.Cmd(nil, PEXPIRE, key, gconv.Int64String(duration.Milliseconds()))); ge.Check(err)
}

func (self RedisClient) PexpireAt(key string, at time.Time) {
	err := self.Core.Do(radix.Cmd(nil, PEXPIREAT, key, gconv.Int64String(at.UnixNano() / int64(time.Millisecond)))); ge.Check(err)
}

func (self RedisClient) Randomkey () (key string) {
	data := radix.MaybeNil{Rcv: &key}
	err := self.Core.Do(radix.Cmd(&data, RANDOMKEY)); ge.Check(err)
	return
}

func (self RedisClient) Rename(oldKey string, newKey string) (err error) {
	return self.Core.Do(radix.Cmd(nil, RENAME, oldKey, newKey))
}
func (self RedisClient) RenameNX(oldKey string, newKey string) (done bool, err error) {
	data := radix.MaybeNil{Rcv: &done}
	err = self.Core.Do(radix.Cmd(&data, RENAMENX, oldKey, newKey))
	return
}