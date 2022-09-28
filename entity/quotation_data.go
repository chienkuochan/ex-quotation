// auth: kunlun
// date: 2019-01-23
// description: 消费MQ消息
package entity

import "time"

type Tick struct {
	Cmd    string     `json:"cmd"`    //指令  tick depth  depth_plus
	Symbol string     `json:"symbol"` //交易对
	Ts     int64      `json:"ts"`     //成交时间戳
	Data   []TickItem `json:"data"`
}

type Depth struct {
	Cmd    string    `json:"cmd"`    //指令  tick depth  depth_plus
	Symbol string    `json:"symbol"` //交易对
	Ts     int64     `json:"ts"`     //成交时间戳
	Data   DepthItem `json:"data"`   //成交数据
}

type DepthItem struct {
	Asks [][]float32
	Bids [][]float32
}

type TickItem struct {
	Amount float32 `json:"amount"` //成交量
	Price  float32 `json:"price"`  //成交价
	Dir    int     `json:"dir"`    //红绿
}

// 模拟Tick数据
func (tick *Tick) MockTick() *Tick {
	var tickData []TickItem
	data := append(tickData, TickItem{Amount: 0.12, Price: 0.11, Dir: 1})
	return &Tick{Cmd: "tick", Symbol: "BTCUSDT", Ts: time.Now().Unix(), Data: data}

}

// 模拟Depth数据
func (depth *Depth) MockDepth() *Depth {
	depthData := DepthItem{Asks: [][]float32{{0.1, 1}, {0.2, 2}}, Bids: [][]float32{{0.3, 0.4}, {0.5, 0.6}}}
	return &Depth{Cmd: "depth", Symbol: "BTCUSDT", Ts: time.Now().Unix(), Data: depthData}
}

// 模拟DepthPlus数据
func (depth *Depth) MockDepthPlus() *Depth {
	depthData := DepthItem{Asks: [][]float32{{0.1, 1}, {0.2, 2}}, Bids: [][]float32{{0.3, 0.4}, {0.5, 0.6}}}
	return &Depth{Cmd: "depthplus", Symbol: "BTCUSDT", Ts: time.Now().Unix(), Data: depthData}
}
