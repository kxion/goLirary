package main

import (
	"goLirary/crypto"
)

func main() {
	crypto.RunMd5("123456")
	crypto.RunSha256("123456")
	crypto.RunSha1("123456")
}
