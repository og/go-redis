package gredis

import (
	"github.com/mediocregopher/radix/v3"
	ge "github.com/og/x/error"
)


const (
	GET = "GET"
	SET = "SET"
	INCR = "INCR"
	DECR = "DECR"
	INCRBY = "INCRBY"
	DECRBY = "DECRBY"
	DEL = "DEL"
	APPEND = "APPEND"
	NX = "NX"
	STRLEN = "STRLEN"
	SETRANGE = "SETRANGE"
	GETRANGE = "GETRANGE"
	INCRBYFLOAT = "INCRBYFLOAT"

	HGET = "HGET"
	HSET = "HSET"
	HSETNX = "HSETNX"
	HINCRBY = "HINCRBY"
	HDEL = "HDEL"
	HLEN = "HLEN"
	HEXISTS = "HEXISTS"
	HINCRBYFLOAT = "HINCRBYFLOAT"


	DUMP = "DUMP"
	EXISTS = "EXISTS"
	PEXPIRE = "PEXPIRE"
	PEXPIREAT = "PEXPIREAT"
	EXPIRE = "EXPIRE"
	EXPIREAT = "EXPIREAT"
	TTL = "TTL"
	PTTL = "PTTL"

	RANDOMKEY = "RANDOMKEY"
	RENAME = "RENAME"
	RENAMENX = "RENAMENX"
)

type RedisClient struct {
	Core *radix.Pool
}

func NewPool(network, addr string, size int, opts ...radix.PoolOpt) (redisClient RedisClient) {
	pool , err:= radix.NewPool(network, addr, size, opts...) ; ge.Check(err)
	redisClient.Core = pool
	return
}
