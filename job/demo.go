package job

import (
	"HereWeGo/initializers"
	"context"
	"github.com/xxl-job/xxl-job-executor-go"
	"log"
)

func init() {
	initializers.RegTask("demoJobHandler", func(cxt context.Context, param *xxl.RunReq) string {
		log.Println("正在执行xxl-job任务")
		return "OK"
	})
}
