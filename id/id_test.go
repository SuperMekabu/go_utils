package id

import (
	"testing"
)

func Test_objHash(t *testing.T) {
	obj := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	obj2 := []byte("!\"#$%&'()~=~|'")
	hash1 := objHash(obj)
	hash2 := objHash(obj)
	hash3 := objHash(obj2)
	t.Run("object hash", func(t *testing.T) {
		if hash1 != hash2 {
			t.Errorf("hash1 %v, hash2 %v", hash1, hash2)
		}

		if hash1 == hash3 {
			t.Errorf("hash1 %v, hash3 %v", hash1, hash3)
		}
	})
}
