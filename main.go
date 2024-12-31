package main

import (
	_ "HereWeGo/consumer"
	_ "HereWeGo/controller"
	"HereWeGo/initializers"
	_ "HereWeGo/job"
	_ "HereWeGo/middlewares"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	r := initializers.Engine
	if err := r.Run(":5050"); err != nil {
		log.Fatalf("端口启动监听失败: %s", err.Error())
	}
}
