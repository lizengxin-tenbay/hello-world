/*
 	web 包：实现了收集数据存入数据库，并通过http请求查询数据库信息
 	@Author: lizengxin
 	@Data:2021/3/30 15:35
 */

package web

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"time"
)

type FxHookUtil struct{}

func NewFxHookUtil2(lc fx.Lifecycle, daemons ...IDaemon) *FxHookUtil {
	checkStart := func(ins IDaemon, errChan chan error) {
		err := ins.Start()
		if err == nil {
			return
		}

		select {
		case errChan <- err:
		default:
		}
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			errChan := make(chan error)
			for _, obj := range daemons {
				go checkStart(obj, errChan)
			}

			select {
			case e := <-errChan:
				return errors.New(`start fail: ` + e.Error())
			case <-time.After(time.Second * 2):
				return nil
			}
		},

		OnStop: func(ctx context.Context) error {
			for _, obj := range daemons {
				obj.Stop()
			}
			time.Sleep(time.Second)
			return nil
		},
	})

	return &FxHookUtil{}
}

type hooker struct{}

func start(_ *hooker) {}

func newHooker(lc fx.Lifecycle, service iWebService, store iMysqlStore, collect iCollect) *hooker {
	NewFxHookUtil2(lc, service, store, collect)
	return &hooker{}
}

func Start() error {
	providers := []interface{}{
		newHooker,
		newWebService,
		newMysqlStore,
		newCollectImpl,
	}

	opts := []fx.Option{fx.Provide(providers...), fx.Invoke(start)}
	if viper.GetInt(`PrintFxEnable`) != 1 {
		opts = append(opts, fx.NopLogger)
	}

	fx.New(opts...).Run()
	return nil
}
