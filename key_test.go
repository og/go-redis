package gredis_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedisClient_Exists(t *testing.T) {
	r.Del("exists")
	assert.Equal(t, r.Exists("exists"), false)
	r.SetString("exists", "abc")
	assert.Equal(t, r.Exists("exists"), true)
}

func TestRedisClient_Expire(t *testing.T) {
	r.Del("Expire")
	r.SetString("Expire", "nimo")
	r.Expire("Expire", 1)
	{
		value , has := r.GetString("Expire")
		assert.Equal(t, value, "nimo")
		assert.Equal(t, has, true)
	}
	time.Sleep(time.Duration(time.Millisecond * 100))
	time.Sleep(time.Duration(time.Second*1))
	{
		value , has := r.GetString("Expire")
		assert.Equal(t, value, "")
		assert.Equal(t, has, false)
	}
}
func TestRedisClient_ExpireAt(t *testing.T) {
	r.Del("ExpireAt")
	r.SetString("ExpireAt", "nimo")
	r.ExpireAt("ExpireAt", time.Now().Add(time.Duration(time.Second*1)))
	{
		value, has := r.GetString("ExpireAt")
		assert.Equal(t, value, "nimo")
		assert.Equal(t, has, true)
	}
	time.Sleep(time.Duration(time.Millisecond*1100))
	{
		value, has := r.GetString("ExpireAt")
		assert.Equal(t, value, "")
		assert.Equal(t, has, false)
	}
}
func TestRedisClient_Pexpire(t *testing.T) {
	r.Del("Pexpire")
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
func TestRedisClient_PexpireAt(t *testing.T) {
	r.Del("PexpireAt")
	r.SetString("PexpireAt", "nimo")
	r.PexpireAt("PexpireAt", time.Now().Add(time.Duration(time.Millisecond*200)))
	{
		value, has := r.GetString("PexpireAt")
		assert.Equal(t, value, "nimo")
		assert.Equal(t, has, true)
	}
	time.Sleep(time.Duration(time.Millisecond*300))
	{
		value, has := r.GetString("PexpireAt")
		assert.Equal(t, value, "")
		assert.Equal(t, has, false)
	}
}

func TestRedisClient_TTL(t *testing.T) {
	r.Del("TTL")
	r.SetString("TTL", "nimo")
	r.Expire("TTL", 2)
	assert.Equal(t, r.TTL("TTL"), 2)
	time.Sleep(time.Duration(time.Second*1))
	assert.Equal(t, r.TTL("TTL"), 1)
}

func TestRedisClient_PTTL(t *testing.T) {
	r.Del("PTTL")
	r.SetString("PTTL", "nimo")
	r.Expire("PTTL", 2)
	{
		value := r.PTTL("PTTL")
		assert.True(t, value< 2000)
		assert.True(t, value> 1900)
	}
	{
		time.Sleep(time.Duration(time.Second*1))
		value := r.PTTL("PTTL")
		assert.True(t, value< 1000)
		assert.True(t, value> 900)
	}
}

func TestRedisClient_Randomkey(t *testing.T) {
	r.SetString("abc", "1")
	r.SetString("efg", "1")
	assert.Equal(t, len(r.Randomkey()) > 0, true)
}
func TestRedisClient_Rename(t *testing.T) {
	r.Del("oldKey")
	r.Del("newKey")
	r.SetString("oldKey", "nimo")
	{
		value, has := r.GetString("newKey")
		assert.Equal(t, value, "")
		assert.Equal(t, has, false)
	}
	err := r.Rename("oldKey", "newKey")
	assert.Equal(t, err, nil)
	{
		value, has := r.GetString("newKey")
		assert.Equal(t, value, "nimo")
		assert.Equal(t, has, true)
	}
	{
		value, has := r.GetString("oldKey")
		assert.Equal(t, value, "")
		assert.Equal(t, has, false)
	}
}

func TestRedisClient_RenamePanic(t *testing.T) {
	r.Del("notExistOldKey")
	err := r.Rename("notExistOldKey", "notExistNewKey")
	assert.Equal(t, err.Error(),"ERR no such key")
	assert.NotEqual(t, err, nil)
}

func TestRedisClient_RenameNX(t *testing.T) {
	r.Del("oldKeyNX")
	r.Del("newKeyNX")
	r.SetString("oldKeyNX", "nimo")
	{
		value, has := r.GetString("newKeyNX")
		assert.Equal(t, value, "")
		assert.Equal(t, has, false)
	}
	done, err := r.RenameNX("oldKeyNX", "newKeyNX")
	assert.Equal(t, err, nil)
	assert.Equal(t, done, true)
	{
		value, has := r.GetString("newKeyNX")
		assert.Equal(t, value, "nimo")
		assert.Equal(t, has, true)
	}
	{
		value, has := r.GetString("oldKeyNX")
		assert.Equal(t, value, "")
		assert.Equal(t, has, false)
	}
	r.SetString("oldKeyNXHas", "nimo")
	r.SetString("newKeyNXHas", "nico")
	{
		done, err := r.RenameNX("oldKeyNXHas", "newKeyNXHas")
		assert.Equal(t, done, false)
		assert.Equal(t, err, nil)
	}

}

func TestRedisClient_RenameNXPanic(t *testing.T) {
	r.Del("notExistOldKey")
	done, err := r.RenameNX("notExistOldKey", "notExistNewKey")
	assert.Equal(t, done, false)
	assert.Equal(t, err.Error(),"ERR no such key")
	assert.NotEqual(t, err, nil)
}