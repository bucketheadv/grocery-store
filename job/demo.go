package job

import (
	"HereWeGo/initializers"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/xxl-job/xxl-job-executor-go"
)

func init() {
	initializers.RegTask("demoJobHandler", func(cxt context.Context, param *xxl.RunReq) string {
		logrus.Info("正在执行xxl-job任务")
		return "OK"
	})
}
