package id

import (
	"crypto/sha256"
	"github.com/google/uuid"
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"time"
)

/*
	Generate non-sortable UUID version 4
	seed is now unix nano
*/
func NewUUID() string {
	t := time.Now()
	obj, err := uuid.NewRandomFromReader(rand.New(rand.NewSource(t.UnixNano())))
	if err != nil {
		log.Fatalf("UUID generate failed: %v", err)
	}
	return obj.String()
}

/*
	Generate non-sortable UUID version 5
	seed is byte slice of some object
*/
func NewUUIDFromObj(obj []byte) string {
	sha1 := uuid.NewSHA1(objHash(obj), obj)
	return sha1.String()
}

func objHash(obj []byte) uuid.UUID {
	hash := sha256.New()
	defer hash.Reset()
	hash.Write(obj)
	bytes := [16]byte{}
	for i, s := range []rune("0123456789ABCDEF") {
		bytes[i] = byte(s)
	}
	return uuid.NewHash(hash, bytes, obj, 5)
}

/*
	Generate sortable ULID
	seed is now unix nano
*/
func NewULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
