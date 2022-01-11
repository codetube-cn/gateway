package bootstrap

import (
	"codetube.cn/gateway/components"
	"log"
)

func Start() {
	//初始化数据库连接
	err := components.DB.MysqlInit()
	if err != nil {
		log.Fatal(err)
	}

	//加载网关
	loadGateway()

	//启动监听服务
	serve()
}
