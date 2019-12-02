package gredis

import (
	"github.com/mediocregopher/radix/v3"
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
	err := self.Core.Do(radix.Cmd(nil, HSET, key, field, value, NX)); ge.Check(err)
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

// func (self RedisClient) StrLen(key string) (value int) {
// 	data := radix.MaybeNil{Rcv: &value}
// 	err := self.Core.Do(radix.Cmd(&data, STRLEN, key))
// 	ge.Check(err)
// 	return
// }
// func (self RedisClient) SetInt(key string, value int) () {
// 	err := self.Core.Do(radix.Cmd(nil, SET, key, gconv.IntString(value))); ge.Check(err)
// }
// func (self RedisClient) SetIntNX(key string, value int) {
// 	err := self.Core.Do(radix.Cmd(nil, SET, key, gconv.IntString(value), NX)); ge.Check(err)
// }
// func (self RedisClient) GetInt(key string) (value int , has bool) {
// 	data := radix.MaybeNil{Rcv: &value}
// 	err := self.Core.Do(radix.Cmd(&data, GET, key))
// 	switch true {
// 	case err != nil:
// 		panic(err)
// 	case data.Nil == true:
// 		return 0, has
// 	case data.Nil == false:
// 		has = true
// 		return
// 	}
// 	return
// }
//
// func (self RedisClient) SetFloat(key string, value float64) () {
// 	stringValue := strconv.FormatFloat(value, 'E', -1, 64)
// 	err := self.Core.Do(radix.Cmd(nil, SET, key, stringValue)); ge.Check(err)
// }
// func (self RedisClient) GetFloat(key string) (value float64 , has bool) {
// 	data := radix.MaybeNil{Rcv: &value}
// 	err := self.Core.Do(radix.Cmd(&data, GET, key))
// 	switch true {
// 	case err != nil:
// 		panic(err)
// 	case data.Nil == true:
// 		return 0, has
// 	case data.Nil == false:
// 		has = true
// 		return
// 	}
// 	return
// }
//
//
// func (self RedisClient) Incr(key string) (value int) {
// 	data := radix.MaybeNil{Rcv: &value}
// 	err := self.Core.Do(radix.Cmd(&data, INCR, key)); ge.Check(err)
// 	return
// }
// func (self RedisClient) Decr(key string)(value int) {
// 	data := radix.MaybeNil{Rcv: &value}
// 	err := self.Core.Do(radix.Cmd(&data, DECR, key)); ge.Check(err)
// 	return
// }
// func (self RedisClient) IncrBy(key string, amount int) (value int) {
// 	data := radix.MaybeNil{Rcv: &value}
// 	err := self.Core.Do(radix.Cmd(&data, INCRBY, key, gconv.IntString(amount))); ge.Check(err)
// 	return
// }
// func (self RedisClient) DecrBy(key string, amount int) (value int) {
// 	data := radix.MaybeNil{Rcv: &value}
// 	err := self.Core.Do(radix.Cmd(&data, DECRBY, key, gconv.IntString(amount))); ge.Check(err)
// 	return
// }
//
// func (self RedisClient) Del(key string) (successDelCount int) {
// 	data := radix.MaybeNil{Rcv: &successDelCount}
// 	err := self.Core.Do(radix.Cmd(&data, DEL, key)); ge.Check(err)
// 	return
// }
//
// func (self RedisClient) Append(key string, value string) () {
// 	err := self.Core.Do(radix.Cmd(nil, APPEND, key, value)); ge.Check(err)
// }
// func (self RedisClient) SetRange(key string, offset int, value string) {
// 	err := self.Core.Do(radix.Cmd(nil, SETRANGE, key, gconv.IntString(offset), value)); ge.Check(err)
// }
// func (self RedisClient) GetRange(key string, start int, end int) (value string) {
// 	data := radix.MaybeNil{Rcv: &value}
// 	err := self.Core.Do(radix.Cmd(&data, GETRANGE, key, gconv.IntString(start), gconv.IntString(end)))
// 	ge.Check(err)
// 	return
// }
// func (self RedisClient) IncrByFloat(key string, amount float64) (value float64) {
// 	data := radix.MaybeNil{Rcv: &value}
// 	stringValue := strconv.FormatFloat(amount, 'E', -1, 64)
// 	err := self.Core.Do(radix.Cmd(&data, INCRBYFLOAT, key, stringValue)); ge.Check(err)
// 	return
// }
//
