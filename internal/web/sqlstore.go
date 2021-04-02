package web

import (
	_ "github.com/go-sql-driver/mysql" //nolint
	"github.com/go-xorm/xorm"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"time"
)

// 用于存储数据库信息
type mysqlStoreOpt struct {
	MysqlURL string // 存储数据库信息
}

// 用于存储数据库信息和xorm.Engine指针
type mysqlStore struct {
	opt mysqlStoreOpt
	*xorm.Engine
}

/*
loadMysqlStoreOpt 赋予具体数据库信息，mysqlStoreOpt为存储信息的结构体
返回值：数据库信息
 */
func loadMysqlStoreOpt() mysqlStoreOpt {
	url := "root:123456@tcp(192.168.108.110:33306)/test?charset=utf8"
	viper.SetDefault(`MYSQL_URL`, url)
	opt := mysqlStoreOpt{
		MysqlURL: viper.GetString(`MYSQL_URL`),
	}
	return opt
}

/*
newXormEngine创建一个xorm.NewEngine实例，url作为数据库类型
返回值：创建成功时返回创建的实例指针，失败时返回error信息
*/
func newXormEngine(url string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		log.Error().Err(err).Str(`url`, url).Msg(`NewEngine`)
		return nil, err
	}

	if err := engine.Ping(); err != nil {
		_ = engine.Close()
		log.Error().Err(err).Msg(`ping`)
		return nil, err
	}

	return engine, nil
}

/*
newMysqlStore 获取数据库信息并创建xorm.Engine
loadMysqlStoreOpt 可获取数据库信息
返回值：数据库信息和xorm.Engine型指针
*/
func newMysqlStore() (iMysqlStore, error) {
	opt := loadMysqlStoreOpt()
	log.Info().Interface(`opt`, &opt).Send()

	engine, err := newXormEngine(opt.MysqlURL)
	if err != nil {
		log.Error().Err(err).Msg("newMysqlStore")
		return nil, err
	}

	log.Info().Msg(`connect mysql ok`)
	return &mysqlStore{
		opt:    opt,
		Engine: engine,
	}, nil
}

func (m mysqlStore) Start() error {
	log.Info().Msg("start")

	for {
		time.Sleep(time.Second)
	}

	return nil
}

func (m mysqlStore) Stop() {
	log.Info().Msg("stop")
	return
}

/*
GetLatest获取数据库最近的n条数据
返回值：从数据库获取的数据，以数组形式返回
 */
func (m *mysqlStore) GetLatest(n int) ([]SysLoad, error) {
	var bean []SysLoad
	err := m.Sql(`select * from sys_load order by created desc limit ?`, n).Find(&bean)

	return bean, err
}

// Update 将数据存储到数据库，load为要更新的数据
func (m *mysqlStore) Update(load SysLoad) error {
	_, err := m.Table(SysLoad{}).InsertOne(load)

	return err
}
