// auth: kunlun
// date: 2019-01-10
// description:
package config

import "testing"

// 执行测试用例需要 将redis.json路径修改为 ../../../redis.json
func TestInitRedisPool(t *testing.T) {
	InitRedisPool()
}
