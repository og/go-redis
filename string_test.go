package gredis_test

import (
	gredis "github.com/og/go-redis"
	grand "github.com/og/x/rand"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)
var r gredis.RedisClient
func init () {
	r = gredis.NewPool("tcp", "127.0.0.1:6379", 10)
}

func TestRedisClient_GetString(t *testing.T) {
	randKey := "GetString:" + grand.StringLetter(10)
	value, has := r.GetString(randKey)
	assert.Equal(t, value, "")
	assert.Equal(t, has, false)
}
func TestRedisClient_SetString(t *testing.T) {
	r.SetString("SetString", "nimo")
	value, has := r.GetString("SetString")
	assert.Equal(t, value, "nimo")
	assert.Equal(t, has, true)
}
func TestRedisClient_SetStringNX(t *testing.T) {
	r.Del("SetStringNX")
	{
		r.SetString("SetStringNX", "1")
		value, has := r.GetString("SetStringNX")
		assert.Equal(t, value, "1")
		assert.Equal(t, has, true)
	}
	{
		value, has := r.GetString("SetStringNX")
		assert.Equal(t, value, "1")
		assert.Equal(t, has, true)
	}
}
func TestRedisClient_SetIntNX(t *testing.T) {
	r.Del("SetIntNX")
	{
		r.SetIntNX("SetIntNX", 1)
		value, has := r.GetInt("SetIntNX")
		assert.Equal(t, value, 1)
		assert.Equal(t, has, true)
	}
	{
		r.SetIntNX("SetIntNX", 2)
		value, has := r.GetInt("SetIntNX")
		assert.Equal(t, value, 1)
		assert.Equal(t, has, true)
	}
}
func TestRedisClient_StrLen(t *testing.T) {
	r.Del("StrLen")
	{
		r.SetString("StrLen", "12345")
		assert.Equal(t, r.StrLen("StrLen"), 5)
	}
	{
		assert.Equal(t, r.StrLen("StrLenNotExist"), 0)
	}
}
func TestRedisClient_Pexpire(t *testing.T) {
	r.SetString("Pexpire", "nimo")
	r.Pexpire("Pexpire", time.Duration(time.Second*1))
	{
		value , has := r.GetString("Pexpire")
		assert.Equal(t, value, "nimo")
		assert.Equal(t, has, true)
	}
	{
		time.Sleep(time.Duration(time.Millisecond * 200))
		time.Sleep(time.Duration(time.Second*1))
		value , has := r.GetString("Pexpire")
		assert.Equal(t, value, "")
		assert.Equal(t, has, false)
	}
}

func TestRedisClient_Incr(t *testing.T) {
	{
		r.Del("Incr")
	}
	{
		value , has := r.GetInt("Incr")
		assert.Equal(t, value, 0)
		assert.Equal(t, has, false)
	}
	{
		assert.Equal(t, r.Incr("Incr"), 1)
	}
	{
		assert.Equal(t, r.Incr("Incr"), 2)
	}
}
func TestRedisClient_IncrBy(t *testing.T) {
	{
		r.Del("IncrBy")
	}
	{
		value , has := r.GetInt("IncrBy")
		assert.Equal(t, value, 0)
		assert.Equal(t, has, false)
	}
	{

		assert.Equal(t, r.IncrBy("IncrBy", 2), 2)
	}
	{
		assert.Equal(t, r.IncrBy("IncrBy",3), 5)
	}
	{
		assert.Equal(t, r.IncrBy("IncrBy",0), 5)
	}
	{
		assert.Equal(t, r.IncrBy("IncrBy",-1), 4)
	}
}
func TestRedisClient_Decr(t *testing.T) {
	{
		r.Del("Decr")
	}
	{
		value , has := r.GetInt("Decr")
		assert.Equal(t, value, 0)
		assert.Equal(t, has, false)
	}
	{
		assert.Equal(t, r.Decr("Decr"), -1)
	}
	{
		assert.Equal(t, r.Decr("Decr"), -2)
	}
}
func TestRedisClient_DecrBy(t *testing.T) {
	{
		r.Del("DecrBy")
	}
	{
		value , has := r.GetInt("DecrBy")
		assert.Equal(t, value, 0)
		assert.Equal(t, has, false)
	}
	{

		assert.Equal(t, r.DecrBy("DecrBy", 2), -2)
	}
	{
		assert.Equal(t, r.DecrBy("DecrBy",3), -5)
	}
	{
		assert.Equal(t, r.DecrBy("DecrBy",0), -5)
	}
	{
		assert.Equal(t, r.DecrBy("DecrBy",-1), -4)
	}
}

func TestRedisClient_Del(t *testing.T) {
	r.SetString("del", "nimo")
	{
		value , has := r.GetString("del")
		assert.Equal(t, value, "nimo")
		assert.Equal(t, has, true)
	}
	{
		assert.Equal(t, r.Del("del"), 1)
		value , has := r.GetString("del")
		assert.Equal(t, value, "")
		assert.Equal(t, has, false)
	}
	{
		assert.Equal(t, r.Del("del"), 0)
	}
}
func TestRedisClient_SetInt(t *testing.T) {
	r.Del("setInt")
	{
		r.SetInt("setInt", 10)
		value , has := r.GetInt("setInt")
		assert.Equal(t, value, 10)
		assert.Equal(t, has, true)
	}
}
func TestRedisClient_GetSetFloat(t *testing.T) {
	r.Del("setFloat")
	{
		var value float64
		value = 1.11
		r.SetFloat("setFloat", value)
		value , has := r.GetFloat("setFloat")
		assert.Equal(t, value, value)
		assert.Equal(t, has, true)
	}
}
func TestRedisClient_Append(t *testing.T) {
	r.Del("append")
	{

		r.Append("append", "1")
		value , has := r.GetString("append")
		assert.Equal(t, value, "1")
		assert.Equal(t, has, true)
	}
	{
		r.Append("append", "2")
		value , has := r.GetString("append")
		assert.Equal(t, value, "12")
		assert.Equal(t, has, true)
	}
}
func TestRedisClient_SetRange(t *testing.T) {
	r.Del("setRange")
	{
		r.SetString("setRange", "abcdef")
		r.SetRange("setRange", 2, "nimo")
		value, has := r.GetString("setRange")
		assert.Equal(t, value, "abnimo")
		assert.Equal(t, has, true)
	}
	{
		r.SetRange("setRangeNotExist", 2, "nimo")
		value, has := r.GetString("setRangeNotExist")
		assert.Equal(t, value, "\x00\x00nimo")
		assert.Equal(t, has, true)
	}
}
func TestRedisClient_GetRange(t *testing.T) {
	r.Del("getRange")
	{
		r.SetString("getRange", "ABCDEF")
		value := r.GetRange("getRange", 2,3)
		assert.Equal(t, "CD", value)
	}
}
func TestRedisClient_GetBool(t *testing.T) {
	r.Del("bool")
	assert.Equal(t, r.GetBool("bool"), false)
	r.SetBool("bool", true)
	assert.Equal(t, r.GetBool("bool"), true)
	r.SetBool("bool", false)
	assert.Equal(t, r.GetBool("bool"), false)
}
func TestRedisClient_IncrByFloat(t *testing.T) {
	r.Del("IncrByFloat")
	assert.Equal(t, r.IncrByFloat("IncrByFloat", 1.1), 1.1)
	assert.Equal(t, r.IncrByFloat("IncrByFloat", 1.1111), 2.2111)
}