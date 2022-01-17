package bootstrap

import (
	"codetube.cn/gateway/components"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//BootErrChan 启动错误
var BootErrChan chan error

func Start() {
	BootErrChan = make(chan error)

	//初始化数据库连接
	err := components.DB.MysqlInit()
	if err != nil {
		log.Fatal(err)
	}

	//从数据库加载网关
	loadGateway()

	//加载路由
	gm, rs, err := getGatewayRoutes()
	if err != nil {
		log.Fatal(err)
	}
	routeGroupsMapping = gm
	gatewayRoutes = rs

	//启动监听服务
	server := httpServer()
	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		log.Println("Gateway Serve: " + server.Addr)
	}()

	//监听重载网关变更
	t := watchGateway()

	//监听退出信号
	go func() {
		sigC := make(chan os.Signal)
		signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)
		BootErrChan <- fmt.Errorf("%", <-sigC)
	}()
	getErr := <-BootErrChan
	log.Println(getErr)

	//关闭重载网关变更
	t.Stop()

	//关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("watch gateway stop ......")
	log.Println("Server Shutdown ......")
}
