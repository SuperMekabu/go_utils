package ids

import (
	"testing"
	"time"
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

func TestNewULID(t *testing.T) {
	now := time.Now()
	entropy := NewEntropy(now)
	var ids []string
	for range []int{1, 2, 3} {
		ids = append(ids, NewULID(entropy))
	}

	first := ids[0]
	last := ids[len(ids)-1]
	t.Run("Generate ULID", func(t *testing.T) {
		if first > last {
			t.Errorf("invalid ULID order, first: %v, last: %v", first, last)
		}

		bef := ids[0]
		for i, s := range ids {
			if i == 0 {
				continue
			}
			if bef > s {
				t.Errorf("invalid ULID order, before: %v, current: %v", bef, s)
			}
			bef = s
		}

		reverse := make([]string, len(ids))
		lastIdx := len(ids) - 1
		for i := 0; i < len(ids)/2+1; i++ {
			reverse[i], reverse[lastIdx-i] = ids[lastIdx-i], ids[i]
		}

		bef = reverse[0]
		for i, s := range reverse {
			if i == 0 {
				continue
			}
			if bef < s {
				t.Errorf("invalid ULID order, before: %v, current: %v", bef, s)
			}
			bef = s
		}

	})
}

func TestNewUUID(t *testing.T) {
	id1 := NewUUID()
	id2 := NewUUID()
	id3 := NewUUID()

	t.Run("UUID v4", func(t *testing.T) {
		if id1 == id2 || id1 == id3 || id2 == id3 {
			t.Errorf("Duplicate uuid, id1: %v, id2: %v, id3: %v", id1, id2, id3)
		}
	})
}

func TestNewUUIDFromObj(t *testing.T) {
	obj := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	obj2 := []byte("!\"#$%&'()~=~|'")

	id1 := NewUUIDFromObj(obj)
	id2 := NewUUIDFromObj(obj)
	id3 := NewUUIDFromObj(obj2)

	t.Run("UUID v5 from byte slice", func(t *testing.T) {
		if id1 != id2 {
			t.Errorf("different UUID from same object, id1: %v, id2: %v", id1, id2)
		}
		if id1 == id3 {
			t.Errorf("same UUID from differ object, id1: %v, id3: %v", id1, id3)
		}
	})
}
