package gluacrypto

import (
	crypto "github.com/cyian-1756/gluacrypto/crypto"
	"github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	L.PreloadModule("crypto", crypto.Loader)
}
