package id

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/google/uuid"
)

// NewUUUIDAsString create deterministic UUUID
func NewUUUIDAsString(input string) string {
	// calculate the MD5 hash of the
	md5hash := md5.New()
	_, err := md5hash.Write([]byte(input))
	if err != nil {
		log.Fatal(err)
	}

	// convert the hash value to a string
	md5string := hex.EncodeToString(md5hash.Sum(nil))

	// generate the UUID from the
	uuid, err := uuid.FromBytes([]byte(md5string[0:16]))
	if err != nil {
		log.Fatal(err)
	}

	return uuid.String()
}
