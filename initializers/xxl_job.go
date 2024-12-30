package initializers

import (
	"context"
	"github.com/xxl-job/xxl-job-executor-go"
	"log"
)

type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	log.Printf("自定义日志 - "+format, a...)
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Fatalf("自定义日志 - "+format, a...)
}

func init() {
	exec := xxl.NewExecutor(
		xxl.ServerAddr("http://localhost:9090"),
		xxl.AccessToken("default_token"),
		xxl.ExecutorPort("9999"),
		xxl.RegistryKey("xxl-job-executor-sample"),
		xxl.SetLogger(&logger{}),
	)

	exec.Init()
	exec.LogHandler(func(req *xxl.LogReq) *xxl.LogRes {
		return &xxl.LogRes{
			Code: 200,
			Msg:  "",
			Content: xxl.LogResContent{
				FromLineNum: req.FromLineNum,
				ToLineNum:   2,
				LogContent:  "这是自定义日志Handler",
				IsEnd:       true,
			},
		}
	})

	exec.RegTask("demoJobHandler", func(cxt context.Context, param *xxl.RunReq) string {
		log.Println("正在执行xxl-job任务")
		return "OK"
	})

	go func() {
		log.Fatal(exec.Run())
	}()
}
