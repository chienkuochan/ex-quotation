// auth: kunlun
// date: 2019-01-10
// description:
package codec

import (
	"common"
)

/*

|    magic   |  len  |  data  |  crc32  |
------------------------------------------
|magicNyxV0.1|   3   |   ABC  | 4354356 |

*/

const (
	QuoteHeader    = "magicNyxV0.1"
	QuoteHeaderLen = 12
)

// 行情编码
func QuoteEncode(message []byte) []byte {
	crcVal := common.Crc32(message)
	return append(
		append(
			append([]byte(QuoteHeader),
				Uint16ToByte(uint16(len(message)))...),
			message...),
		Uint32ToByte(crcVal)...,
	)
}

//行情解码
func QuoteDecode(buffer []byte, ch chan []byte) []byte {
	var i int
	len := len(buffer)
	var messageLen uint16
	for i = 0; i < len; i++ {
		if len < int(QuoteHeaderLen) {
			break
		}
		if string(buffer[i:QuoteHeaderLen]) == QuoteHeader {
			messageLen = ByteToUint16(buffer[i+QuoteHeaderLen : i+QuoteHeaderLen+2])
			if len <= i+QuoteHeaderLen+2+int(messageLen)+4 {
				break
			}
			data := buffer[i+QuoteHeaderLen : i+QuoteHeaderLen+int(messageLen)]
			ch <- data
			i += QuoteHeaderLen + int(messageLen) - 1
		}
	}
	if i == len {
		return make([]byte, 0)
	}
	return buffer[i+QuoteHeaderLen+2 : i+QuoteHeaderLen+2+int(messageLen)]
}
