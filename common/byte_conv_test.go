// auth: kunlun
// date: 2019-01-10
// description:
package common

import "testing"

func TestUint32ToByte(t *testing.T) {
	Uint32ToByte(13)
}

func BenchmarkUint32ToByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32ToByte(12)
	}
}
