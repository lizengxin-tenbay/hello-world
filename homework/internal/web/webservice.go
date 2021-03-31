package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type webServerOpt struct {
	ListenAdrr string		// 监听端口
}

// 存储数据库接口
type webServiceImpl struct {
	store 	iMysqlStore		// 数据库接口
}

func (w *webServiceImpl) Stop() {
	log.Info().Msg("stop")
}

//  loadWebServerOpt 设置具体监听端口，webServerOpt用于存储监听端口值
func loadWebServerOpt() webServerOpt {
	viper.SetDefault(`ListenAdrr`, ":8888")

	var opt = webServerOpt{
		ListenAdrr: viper.GetString(`ListenAdrr`),
	}

	return opt
}

// newGinEngine 创建gin.engine实例，并返回实例指针
func newGinEngine() *gin.Engine {
	engine := gin.Default()
	return engine
}

// newWebService 获取数据库接口，并返回接口
func newWebService(store iMysqlStore) (iWebService, error) {
	return &webServiceImpl{
		store: store,
	}, nil
}

func (w *webServiceImpl) Start() error {
	log.Info().Msg("Start")

	engine := newGinEngine()
	engine.GET("/get", w.handleGet)

	opt := loadWebServerOpt()
	err := engine.Run(opt.ListenAdrr)
	if err != nil {
		log.Error().Err(err).Msg("engine.Run() fail")
		return err
	}

	return nil
}

/*
handleGet 处理http请求
 */
func (w *webServiceImpl) handleGet(c *gin.Context) {
	x, err := w.store.GetLatest(10)	// 用于获取最新的数据，c.JSON将数据输出到端口
	if err != nil {
		c.JSON(200, gin.H{"error":err})
		return
	}

	c.JSON(200, gin.H{"result":x})
}
