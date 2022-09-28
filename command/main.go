// auth: kunlun
// date: 2019-01-14
// description:
package main

import (
	"fmt"
	"net/http"
	"websocket"
)

func main() {

	//server.Start()

	http.HandleFunc("/", websocket.Handle)
	go websocket.Ping()
	go websocket.Task()

	http.ListenAndServe("0.0.0.0:1234", nil)

	slice1 := []string{"1", "2"}
	slice2 := []string{"3"}
	fmt.Println(append(slice1, slice2...))

	//var wg sync.WaitGroup
	//quit := make(chan int)
	//wg.Add(1)
	//ticker1 := time.NewTicker(1 * time.Second)
	//count := 1
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("child goroutine bootstrap start")
	//	for {
	//		select {
	//		case <-ticker1.C:
	//			count++
	//			fmt.Printf("ticker: %d\n", count)
	//		case <-quit:
	//			fmt.Println("work well")
	//			ticker1.Stop()
	//			return
	//		}
	//	}
	//	fmt.Println("child goroutine bootstrap end")
	//}()
	//time.Sleep(50 * time.Second)
	//quit <- 1
	//wg.Wait()
}
