// auth: kunlun
// date: 2019-01-10
// description:
package config

//type Config struct {
//	Redis    *Redis
//	RabbitMQ *RabbitMQ
//	Mysql    *Mysql
//}

type RabbitMQ struct {
	Host     string
	Port     string
	Vhost    string
	UserName string
	Password string
	Exchange []*Exchange
}

type Mysql struct {
	UserName string
	Password string
	Url      string
}
