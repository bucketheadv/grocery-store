package main

import (
	"fmt"
	"github.com/bucketheadv/infra-core/modules/logger"
	"go.uber.org/zap"
	_ "grocery-store/api"
	"grocery-store/conf"
	_ "grocery-store/consumer"
	"grocery-store/initial"
	_ "grocery-store/job"
)

func main() {
	r := initial.Engine
	logger.InitWithConfig(logger.Config{
		InfoLogPath:  "log/info.log",
		ErrorLogPath: "log/error.log",
		Debug:        true,
		Level:        int8(zap.DebugLevel),
		Rotate: logger.RotateCfg{
			MaxSize:    1024,
			MaxAge:     7,
			MaxBackups: 30,
			Compress:   true,
		},
	})
	var port = fmt.Sprintf(":%d", conf.Config.Server.Port)
	if err := r.Run(port); err != nil {
		logger.Fatalf("端口启动监听失败: %s", err.Error())
	}
}
