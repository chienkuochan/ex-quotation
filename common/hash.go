// auth: kunlun
// date: 2019-01-10
// description:
package common

import (
	"crypto/md5"
	"hash/crc32"
)

const (
	Md5Salt = "MD5Salt"
)

//get crc32 hash
func Crc32(bytes []byte) uint32 {
	return crc32.ChecksumIEEE(bytes)
}

//check crc32 sum
func CheckCrc(src, dest uint32) bool {
	if src == dest {
		return true
	}
	return false
}

//md5
func Md5Byte(bytes []byte) [16]byte {
	bytes = append([]byte(Md5Salt))
	return md5.Sum(bytes)
}

//md5
func Md5String(message string) [16]byte {
	message = message + Md5Salt
	return md5.Sum([]byte(message))
}
