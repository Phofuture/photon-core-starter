package photonCoreStarter

import (
	"context"
	"github.com/dennesshen/photon-core-starter/bean"
	"github.com/dennesshen/photon-core-starter/configuration"
	"github.com/dennesshen/photon-core-starter/core"
	"github.com/dennesshen/photon-core-starter/log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type ContextSendSign struct{}

func Run() {
	mainContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 讀取配置檔
	configuration.InitConfiguration()
	
	// 啟動Bean容器初始化
	bean.StartBeanManagement()
	
	// 核心依賴初始化
	for _, action := range core.GetCoreDependencies() {
		if err := action(mainContext); err != nil {
			slog.Error("init core dependencies error", "error", err)
			return
		}
	}
	// 核心包初始化配置
	_ = log.StartLogger()
	
	log.Logger().Info(mainContext, "application is starting")
	
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	once := sync.Once{}
	f := func() {
		once.Do(func() {
			sign <- syscall.SIGTERM
			cancel()
		})
	}
	
	// 附加模組初始化
	startAddModule(mainContext, f)
	
	// 專案初始化
	select {
	case <-mainContext.Done():
	default:
		startProject(mainContext, f)
		log.Logger().Info(mainContext, "application is running")
	}
	
	<-sign
	log.Logger().Info(mainContext, "application is stopping")
	
	// 附加模組關閉
	for _, action := range core.GetShutdownAddModule() {
		_ = action(context.Background())
	}
	
	// 核心包關閉
	_ = log.ShutdownLogger()
	
	// 關閉核心依賴
	for _, action := range core.GetShutdownCoreDependencies() {
		func() {
			thisCtx, thisCancel := context.WithTimeout(mainContext, 10*time.Second)
			defer thisCancel()
			_ = action(thisCtx)
		}()
	}
}

func startAddModule(main context.Context, sendSign func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(core.GetAddModule()))
	for _, action := range core.GetAddModule() {
		go func() {
			childCtx := context.WithValue(main, ContextSendSign{}, sendSign)
			defer func() {
				wg.Done()
				if r := recover(); r != nil {
					log.Logger().Error(childCtx, "process panic,", "reason", r)
					sendSign()
				}
			}()
			if err := action(childCtx); err != nil {
				log.Logger().Error(childCtx, "init action error,", "error", err)
				sendSign()
				return
			}
		}()
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-main.Done():
	case <-done:
	}
}

func startProject(ctx context.Context, sendSign func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Logger().Error(ctx, "panic:", "error", r)
			sendSign()
		}
	}()
	for _, action := range core.GetProjectInit() {
		if err := action(ctx); err != nil {
			log.Logger().Error(ctx, "init project error", "error", err)
			sendSign()
		}
	}
}
