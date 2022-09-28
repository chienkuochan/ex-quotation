// auth: kunlun
// date: 2019-01-23
// description: 存储数据结构
package entity

type StoreData struct {
	Id          int64   `json:"id"`           //主键
	Symbol      string  `json:"symbol"`       //交易对
	SymbolAlias string  `json:"symbol_alias"` //交易对别名
	Dir         int     `json:"dir"`          //交易方向显示红绿
	Ts          int64   `json:"ts"`           //成交时间 纳秒
	High        float32 `json:"high"`         //最高
	Low         float32 `json:"low"`          //最低
	Amount      float32 `json:"amount"`       //成交量
	Vol         float32 `json:"vol"`          //成交额
	Open        float32 `json:"open"`         //今开价
	Close       float32 `json:"close"`        //昨收价
	CreateTime  int64   `json:"create_time"`  //创建时间
	UpdateTime  int64   `json:"update_time"`  //更新时间
}
