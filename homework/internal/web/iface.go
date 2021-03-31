package web

type IDaemon interface {
	Start() error
	Stop()
}

// 网络服务接口
type iWebService interface {
	IDaemon
}

// 数据存储接口
type iMysqlStore interface {
	IDaemon
	GetLatest(n int) ([]SysLoad, error)  	// 获取数据库最新n条数据
	Update(SysLoad) error   				// 更新数据库数据
}

// 数据搜集接口
type iCollect interface {
	IDaemon
}
