// auth: kunlun
// date: 2019-01-10
// description:
package common

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// 字节转换 int  使用大端模式
//
// 大端  高位在内存低地址   低位在内存高地址
func BytesToInt(b []byte) int {
	byteBuffer := bytes.NewBuffer(b)
	var value int32
	binary.Read(byteBuffer, binary.BigEndian, &value)
	return int(value)
}

// 字节转 uint16
func ByteToUint16(b []byte) uint16 {
	var val uint16 = 0
	for i := 0; i < len(b); i++ {
		val = val + uint16(uint(b[i])<<uint(8*i))
	}
	return val
}

// int 转字节
func IntToBytes(val int) []byte {
	value := int32(val)
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, value)
	return byteBuffer.Bytes()
}

// 字节转 16进制
func ByteToHex(buffer []byte) string {
	return hex.EncodeToString(buffer)
}

// uint16 to byte
func Uint16ToByte(val uint16) []byte {
	return []byte{byte(val), byte(val >> 8)}

}

// uint32 to byte
func Uint32ToByte(val uint32) []byte {
	return []byte{byte(val), byte(val >> 8), byte(val >> 16), byte(val >> 24)}
}

// hex string to byte
func HexToByte(str string) []byte {
	value, _ := hex.DecodeString(str)
	return value
}

// byte to string
func ByteToString(b *[]byte) *string {
	str := bytes.NewBuffer(*b)
	result := str.String()
	return &result
}
