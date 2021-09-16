package shadowsocks

import (
	"crypto/md5"

	"github.com/riobard/go-shadowsocks2/shadowaead"
)

// key-derivation function from original Shadowsocks
func ShadowsocksKDF(password string, keyLen int) []byte {
	var b, prev []byte
	h := md5.New()
	for len(b) < keyLen {
		h.Write(prev)
		h.Write([]byte(password))
		b = h.Sum(b)
		prev = b[len(b)-h.Size():]
		h.Reset()
	}
	return b[:keyLen]
}

func TryCipherWork(password string, salt []byte) bool {
	key := ShadowsocksKDF(password, 32) // temp for test
	metaCipher, err := shadowaead.Chacha20Poly1305(key)
	if err != nil {
		println(err.Error())
		return false
	}
	_, err = metaCipher.Decrypter(salt)
	if err != nil {
		println(err.Error())
		return false
	}
	return true
}
