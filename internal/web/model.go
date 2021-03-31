package web

import "time"

// 数据库sys_load对应的结构体数据
type SysLoad struct {
	Id 			int64  		`json:"id" xorm:"pk notnull autoincr"`
	Load1m  	int    		`json:"load1m" xorm:"load1m"`	//
	Load5m  	int    		`json:"load5m" xorm:"load5m"`	//
	Load15m  	int    		`json:"load15m" xorm:"load15m"`
	MemUsed  	int    		`json:"memused" xorm:"memused"`
	MemTotal  	int    		`json:"memtotal" xorm:"memtotal"`
	Created  	time.Time 	`json:"created,omitempty" xorm:"created"`
}
