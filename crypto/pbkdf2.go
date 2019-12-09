package gluacrypto_crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	lua "github.com/yuin/gopher-lua"
	"golang.org/x/crypto/pbkdf2"
)

func pbkdf2Fn(L *lua.LState) int {
	passphrase := lua.LVAsString(L.Get(1))
	salt := []byte(lua.LVAsString(L.Get(2)))
	iter := L.ToInt(3)
	keyLen := L.ToInt(4)
	algorithm := lua.LVAsString(L.Get(5))
	var h []byte
	switch algorithm {
	case "md5":
		h = pbkdf2.Key([]byte(passphrase), salt, iter, keyLen, md5.New)
	case "sha1":
		h = pbkdf2.Key([]byte(passphrase), salt, iter, keyLen, sha1.New)
	case "sha256":
		h = pbkdf2.Key([]byte(passphrase), salt, iter, keyLen, sha256.New)
	case "sha512":
		h = pbkdf2.Key([]byte(passphrase), salt, iter, keyLen, sha512.New)
	default:
		L.Push(lua.LNil)
		L.Push(lua.LString("unsupported algorithm"))
		return 2
	}
	L.Push(lua.LString(h))
	return 1

}
