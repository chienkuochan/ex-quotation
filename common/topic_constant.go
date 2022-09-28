// auth: kunlun
// date: 2019-01-16
// description:
package common

const (
	Kline          = "kline" //K线
	UnAck          = 3       //服务端检测到客户端未回包次数
	PingInterval   = 1       //服务端主动向客户端发送心跳频率
	DefaultRate    = 1       //默认推送频率
	Sub            = "sub"   //订阅
	UnSub          = "unsub" //取消订阅
	Ping           = "ping"  //心跳发送方
	Pong           = "pong"  //心跳接收方
	OneMinuter     = 60      //分钟
	FiveMinuter    = 300     //5分钟
	FifteenMinuter = 900     //15分钟
	ThrityMinuter  = 1800    //30分钟
	SixtyMinuter   = 3600    //60分钟
	FourHour       = 14400   //4小时
	OneDay         = 172800  //1天
	OneMonth       = 5184000 //1月
	OneWeek        = 1209600 //1周
)
