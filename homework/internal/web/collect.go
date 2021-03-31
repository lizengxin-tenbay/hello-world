package web

import (
	"github.com/rs/zerolog/log"
	"time"
)

type collectOpt struct {}

type collectImpl struct {
	opt 	collectOpt
	store 	iMysqlStore
}

func loadCollectOpt() collectOpt {
	opt := collectOpt{}
	return opt
}

func newCollectImpl(store iMysqlStore) (iCollect, error) {
	opt := loadCollectOpt()
	return &collectImpl{
		opt: opt,
		store: store,
	}, nil
}

/*
Start 将具体数据更新到数据库
store.Update() 用于具体执行将数据更新到数据库
 */
func (c *collectImpl) Start() error {
	log.Info().Msg("start")

	for i := 0; i < 10; i++{
		time.Sleep(time.Second*10)
		load := SysLoad{
			Load1m: 3,
			Load5m: 6,
			Load15m: 9,
			MemUsed: 10,
			MemTotal: 20,
			Created: time.Now(),
		}

		err := c.store.Update(load)
		if err != nil {
			log.Error().Err(err).Msg(`Update`)
			return err
		}
	}

	return nil
}

func (c *collectImpl) Stop() {
	log.Info().Msg("stop")
}
