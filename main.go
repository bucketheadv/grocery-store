package main

import (
	"fmt"
	"github.com/bucketheadv/infra-gin"
	"github.com/sirupsen/logrus"
	_ "grocery-store/api"
	"grocery-store/conf"
	_ "grocery-store/consumer"
	_ "grocery-store/job"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	r := infra_gin.Engine
	var port = fmt.Sprintf(":%d", conf.Config.Server.Port)
	if err := r.Run(port); err != nil {
		logrus.Fatalf("端口启动监听失败: %s", err.Error())
	}
}
