package gredis

import (
	"github.com/mediocregopher/radix/v3"
	gconv "github.com/og/x/conv"
	ge "github.com/og/x/error"
)

func (self RedisClient) HGetBool(key string, field string) (value bool) {
	data := radix.MaybeNil{Rcv: &value}
	err := self.Core.Do(radix.Cmd(&data, HGET, key, field))
	switch true {
	case err != nil:
		panic(err)
	case data.Nil == true:
		return false
	}
	return
}
func (self RedisClient) HSetBool(key string, field string, value bool) {
	valueString := ""
	if value {
		valueString = "1"
	} else {
		valueString = "0"
	}
	err := self.Core.Do(radix.Cmd(nil, HSET, key, field, valueString)); ge.Check(err)
}


func (self RedisClient) HSetString(key string, field string, value string) {
	err := self.Core.Do(radix.Cmd(nil, HSET, key, field, value)); ge.Check(err)
}
func (self RedisClient) HSetStringNX(key string, field string, value string) {
	err := self.Core.Do(radix.Cmd(nil, HSETNX, key, field, value)); ge.Check(err)
}
func (self RedisClient) HGetString(key string, field string) (value string, has bool) {
	data := radix.MaybeNil{Rcv: &value}
	err := self.Core.Do(radix.Cmd(&data, HGET, key, field))
	switch true {
	case err != nil:
		panic(err)
	case data.Nil == true:
		return "", has
	case data.Nil == false:
		has = true
		return
	}
	return
}

func (self RedisClient) HSetInt(key string, field string, value int) {
	err := self.Core.Do(radix.Cmd(nil, HSET, key, field, gconv.IntString(value))); ge.Check(err)
}
func (self RedisClient) HSetIntNX(key string, field string, value int) {
	err := self.Core.Do(radix.Cmd(nil, HSETNX, key, field, gconv.IntString(value))); ge.Check(err)
}
func (self RedisClient) HGetInt(key string, field string) (value int, has bool) {
	data := radix.MaybeNil{Rcv: &value}
	err := self.Core.Do(radix.Cmd(&data, HGET, key, field))
	switch true {
	case err != nil:
		panic(err)
	case data.Nil == true:
		return 0, has
	case data.Nil == false:
		has = true
		return
	}
	return
}

func (self RedisClient) HLen(key string) (value int) {
	data := radix.MaybeNil{Rcv: &value}
	err := self.Core.Do(radix.Cmd(&data, HLEN, key))
	ge.Check(err)
	return
}
func (self RedisClient) HDel (key ...string) ( deletedCount int){
	data := radix.MaybeNil{Rcv: &deletedCount}
	err := self.Core.Do(radix.Cmd(&data, HLEN, key...))
	ge.Check(err)
	return
}
func (self RedisClient) HExists(key string, field string) (exists bool) {
	data := radix.MaybeNil{Rcv: &exists}
	err := self.Core.Do(radix.Cmd(&data, HEXISTS, key, field))
	ge.Check(err)
	return
}