package main

import (
	"fmt"
	"crypto/sha1"
	"hash/crc32"
)

func main() {
	crc := crc32.NewIEEE()
	crc.Write([]byte("test"))
	val := crc.Sum32()
	fmt.Println(val)


	sha := sha1.New()
	sha.Write([]byte("test"))
	bs := sha.Sum([]byte{})
	fmt.Println(bs)
}
