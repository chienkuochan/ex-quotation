// auth: kunlun
// date: 2019-01-18
// description:
package market

import "time"

// 大盘基本数据
type MarketInfo struct {
	Id         int64     `json:"id"`          //主键
	Name       string    `json:"name"`        //大盘名称
	Status     int       `json:"status"`      //状态 0 收盘  1开盘
	CreateTime time.Time `json:"create_time"` //创建时间
	UpdateTime time.Time `json:"update_time"` //更新时间
}

// 大盘日成交行情
type MarketDay struct {
	Id           int64     `json:"id"`            //主键
	MarketId     int64     `json:"market_id"`     //大盘ID
	OpenIndex    float32   `json:"open_index"`    //今开盘指数
	CloseIndex   float32   `json:"close_index"`   //昨日收盘指数
	High         float32   `json:"high"`          //最高指数
	Low          float32   `json:"low"`           //最低指数
	Amount       float32   `json:"amount"`        //成交量
	Vol          float32   `json:"vol"`           //总成交额
	CurrentIndex float32   `json:"current_index"` //当前指数
	CreateTime   time.Time `json:"create_time"`   //创建时间
}

// 大盘实时成交行情
type MarketDealQuotation struct {
	Id           int64     `json:"id"`            //主键
	MarketId     int64     `json:"market_id"`     //大盘id
	OpenIndex    float32   `json:"open_index"`    //今开盘指数
	CloseIndex   float32   `json:"close_index"`   //昨收盘指数
	High         float32   `json:"high"`          //最高指数
	Low          float32   `json:"low"`           //最低指数
	Amount       float32   `json:"amount"`        //成交总量
	Vol          float32   `json:"vol"`           //成交总额
	CurrentRatio float32   `json:"current_ratio"` //当前指数
	Ts           int64     `json:"ts"`            //成交时间戳
	CreateTime   time.Time `json:"create_time"`   //创建时间
}

// 大盘规则
type MarketRule struct {
	Id             int64     `json:"id"`              //主键
	MarketId       int64     `json:"market_id"`       //大盘id
	StockAmplitude int       `json:"stock_amplitude"` //大盘振幅
	OpenTime       time.Time `json:open_time`         //开盘时间
	CloseTime      time.Time `json:"close_time"`      //收盘时间
	CreateTime     time.Time `json:"create_time"`     //创建时间
	UpdateTime     time.Time `json:"update_time"`     //更新时间
}
