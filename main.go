package main

import (
	_ "HereWeGo/api"
	"HereWeGo/components"
	_ "HereWeGo/consumer"
	_ "HereWeGo/job"
	_ "HereWeGo/middlewares"
	"github.com/sirupsen/logrus"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	r := components.Engine
	if err := r.Run(":5050"); err != nil {
		logrus.Fatalf("端口启动监听失败: %s", err.Error())
	}
}
