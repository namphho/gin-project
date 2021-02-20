package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type Hasher interface {
	Hash(data string) string
}

type md5Hash struct {}

func NewMd5Hash()  *md5Hash{
	return &md5Hash{}
}

func (receiver *md5Hash) Hash(data string) string{
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
