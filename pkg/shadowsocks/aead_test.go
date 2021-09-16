package shadowsocks

import (
	"crypto/rand"
	"testing"

	"github.com/riobard/go-shadowsocks2/shadowaead"
)

func TestTryCipherWork(t *testing.T) {

	plaintext := []byte("hello world")
	res := make([]byte, len(plaintext))

	passwordList := []string{"password1", "password2", "password3"}

	for _, password := range passwordList {
		key := ShadowsocksKDF(password, 32)
		metaCipher, err := shadowaead.Chacha20Poly1305(key)
		if err != nil {
			t.Fatalf(err.Error())
		}
		salt := make([]byte, metaCipher.SaltSize())
		if _, err := rand.Read(salt); err != nil {
			t.Fatalf(err.Error())
		}
		enc, _ := metaCipher.Encrypter(salt)

		enc.Seal(res, plaintext, nil, nil)

	}

}
