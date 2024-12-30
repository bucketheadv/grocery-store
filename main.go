package main

import (
	"HereWeGo/config"
	"HereWeGo/middlewares"
	"github.com/gin-gonic/gin"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	r := gin.Default()
	middlewares.Load(r)
	config.InitRouter(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("端口启动监听失败: %s", err.Error())
	}
}
