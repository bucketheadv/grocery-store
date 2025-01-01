package job

import (
	"HereWeGo/components"
	"context"
	"github.com/shopspring/decimal"
	"github.com/xxl-job/xxl-job-executor-go"
)

func init() {
	components.RegTask("demoJobHandler", func(cxt context.Context, param *xxl.RunReq) string {
		components.LogJobInfo(param, "正在执行xxl-job任务")
		data, _ := decimal.NewFromString("0.01")
		components.LogJobInfo(param, "BigDecimal数据: %s", data)
		return "OK"
	})
}
