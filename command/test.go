// auth: kunlun
// date: 2019-01-22
// description:
package main

import (
	"common"
	"fmt"
)

func main() {

	//node, err := common.NewNode(1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//ch := make(chan common.ID)
	//count := 100000
	//for i := 0; i < count; i++ {
	//	go func() {
	//		id := node.Generate()
	//		ch <- id
	//	}()
	//}
	//
	//defer close(ch)
	//
	//m := make(map[common.ID]int)
	//for i := 0; i < count; i++ {
	//	id := <-ch
	//	fmt.Println(id)
	//	_, ok := m[id]
	//	if ok {
	//		fmt.Println("ID is not unique")
	//		return
	//	}
	//	m[id] = i
	//}
	//fmt.Println("All", count, "success")
	value, err := common.RedisGetString("nnn")
	fmt.Println(value, err)

}
