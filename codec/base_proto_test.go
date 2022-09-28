// auth: kunlun
// date: 2019-01-11
// description:
package codec

import "testing"

func BenchmarkNewDefaultProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDefaultProto("发送信息")
	}
}
