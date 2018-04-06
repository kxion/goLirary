package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func RunMd5(key string) {
	h := md5.New()
	io.WriteString(h, key)
	pwd := h.Sum(nil)
	fmt.Println(hex.EncodeToString(pwd))
}

//hash的长度不一样
func RunSha256(key string) {
	h := sha256.New()
	io.WriteString(h, key)
	pwd := h.Sum(nil)
	fmt.Println(hex.EncodeToString(pwd))
}

func RunSha1(key string) {
	h := sha1.New()
	io.WriteString(h, key)
	pwd := h.Sum(nil)
	fmt.Println(hex.EncodeToString(pwd))
}
