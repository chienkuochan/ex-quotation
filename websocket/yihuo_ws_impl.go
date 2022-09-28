// auth: kunlun
// date: 2019-01-16
// description:
package websocket

import (
	"common"
	"encoding/json"
	"entity"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"handle"
	"log"
	"request"
	"strings"
	"sync"
	"time"
)

type Connection struct {
	wsConnect *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	isClosed  bool
}

var GlobalOption = make(map[string]SubOption)

// 订阅参数选项
type SubOption struct {
	Conn         *Connection       //连接
	LastPushTime time.Time         //最后推送时间
	Rate         int               //推送速率  默认5秒一次
	Symbol       map[string]string //交易对
	Type         string            //推送类型
	UnAckCounter int               //未响应计数器
}

// 初始化连接
func InitConnection(wsCon *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConnect: wsCon,
		inChan:    make(chan []byte, 1024),
		outChan:   make(chan []byte, 1024),
		closeChan: make(chan byte, 1)}

	// 建立连接
	addr := wsCon.RemoteAddr().String()
	subOption := SubOption{Conn: conn}
	GlobalOption[addr] = subOption
	go conn.readLoop()
	go conn.writeLoop()
	return
}

// 读取消息
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
		result, err := MessageHandle(conn, data)
		if err != nil {
			log.Panic(err)
		}
		v, err1 := json.Marshal(result)
		if err1 != nil {
			log.Panic(err1)
		}
		conn.WriteMessage(v)
	case <-conn.closeChan:
		err = errors.New("connection is closed")
		fmt.Println("Connection is closed")
	}
	return
}

// 写消息
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
		fmt.Println("Connection is closed")
	}
	return
}

// 关闭连接
func (conn *Connection) Close() {
	conn.wsConnect.Close()
	conn.mutex.Lock()
	//客户端下线,删除链接
	addr := conn.wsConnect.RemoteAddr().String()
	delete(GlobalOption, addr)
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// 循环读取
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)

	for {
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			goto ERR
		}
	}

ERR:
	conn.Close()
}

// 循环写入
func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

// 消息处理器
// 订阅 & 取消订阅 消息解析
func MessageHandle(conn *Connection, data []byte) (result *request.SubResult, err error) {
	result = &request.SubResult{}
	var sub request.SubRequest

	err = json.Unmarshal(data, &sub)
	if err != nil {
		return result.SubFailure("100010", "订阅错误"), err
	}

	if sub.OpK == common.Sub {
		str := strings.Split(sub.Opv, ".")
		Sub(conn.wsConnect, str[1], "", sub.Rate)
		return result.SubSuccess(), nil
	} else if sub.OpK == common.UnSub {
		Unsub(conn, sub)
	}
	return
}

// 定时任务时间间隔为1秒
// 遍历循环当前连接Map
// 检查每个连接的订阅 交易对  推送频率
// 按需推送
func Task() {
	task := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-task.C:
			for key, value := range GlobalOption {
				var rate int
				if value.Rate == 0 {
					rate = common.DefaultRate
				} else {
					rate = value.Rate
				}
				symbols := value.Symbol
				if len(symbols) != 0 {
					//TODO 区分订阅为 1min  5min  15min

					if int(time.Now().Unix()-value.LastPushTime.Unix()) >= rate {
						fmt.Printf("Time Is Coming，Starting Push : %s\n", value.Conn.wsConnect.RemoteAddr().String())
						result := request.NewTick()
						data, err := json.Marshal(result)
						if err != nil {
							panic(err)
						}
						value.Conn.WriteMessage(data)
						value.LastPushTime = time.Now()
						GlobalOption[key] = value
					}
				}
			}
		}
	}
}

// 订阅
// 订阅信息,当前连接全局记录
func Sub(conn *websocket.Conn, symbol, types string, rate int) {
	address := conn.RemoteAddr().String()
	if _, ok := GlobalOption[address]; ok {
		srcOp := GlobalOption[address]
		if _, ok := srcOp.Symbol[symbol]; !ok {
			srcOp.Symbol[symbol] = symbol
			srcOp.Type = types
			srcOp.Rate = rate
			srcOp.LastPushTime = time.Now()
			fmt.Printf("Subscribe, Client Address: %s", address)
		} else {
			fmt.Printf("Subscribe Repeat,Client Address: %s", address)
			srcOp.Rate = rate
			srcOp.Type = types
		}
		GlobalOption[address] = srcOp

	}

}

// 取消订阅
// all 全部取消,连接还存在
// [btc,eth] 指定交易对取消 连接还存在
// 返回 SubResult
func Unsub(conn *Connection, subRequest request.SubRequest) (result *request.SubResult) {
	opv := subRequest.Opv
	addr := conn.wsConnect.RemoteAddr().String()
	//首先判断次那个月信息是否存在
	if _, ok := GlobalOption[addr]; !ok {
		fmt.Printf("Not Find Subscribe Infomation,Client Address: %s", addr)
		return result.UnSubFailure("10006", "Please make sure Already Subscribe")
	}

	option := GlobalOption[addr]
	if opv == "all" {
		option.Symbol = make(map[string]string)
	} else {
		opvSlice := strings.Split(opv, ",")
		for _, value := range opvSlice {
			delete(option.Symbol, option.Symbol[value])
		}
	}
	return result.UnSubSuccess()
}

// 服务器主动向客户端发送心跳包
// 频率为 5s 一次
// 客户端连续忽略2次,服务端主动断开连接
func Ping() {
	task := time.NewTicker(common.PingInterval * time.Millisecond)
	for {
		select {
		case <-task.C:

			//TODO 模拟发送数据
			tick := entity.Tick{}
			val := tick.MockTick()
			data, _ := json.Marshal(val)
			handle.PublishMessage("go.direct.exchange", "bind1", data)

			for k, v := range GlobalOption {
				//fmt.Println("心跳检测")
				unAck := v.UnAckCounter
				if unAck >= common.UnAck {
					//fmt.Printf("\r 客户端未响应，主动断开 %d",time.Now().UnixNano())
					//v.Conn.Close()
					//continue
				}
				conn := v.Conn
				ping := common.NewPing()
				data, err := json.Marshal(ping)
				if err != nil {
					panic(err)
				}
				conn.WriteMessage(data)
				//未响应计数器 +1
				v.UnAckCounter = unAck + 1
				GlobalOption[k] = v
			}
		}
	}
}
