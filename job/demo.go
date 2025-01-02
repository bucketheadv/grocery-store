package job

import (
	"HereWeGo/initializer"
	"context"
	"github.com/shopspring/decimal"
	"github.com/xxl-job/xxl-job-executor-go"
)

func init() {
	jobClient := initializer.XxlJobClient
	jobClient.RegTask("demoJobHandler", func(cxt context.Context, param *xxl.RunReq) string {
		jobClient.LogJobInfo(param, "正在执行xxl-job任务")
		data, _ := decimal.NewFromString("0.01")
		jobClient.LogJobInfo(param, "BigDecimal数据: %s", data)
		return "OK"
	})
}
