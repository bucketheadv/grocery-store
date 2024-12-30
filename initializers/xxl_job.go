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
	log.Printf("自定义日志 - "+format, a...)
}

func init() {
	conf := GetConfig().XxlJob
	exec := xxl.NewExecutor(
		xxl.ServerAddr(conf.ServerAddr),
		xxl.AccessToken(conf.AccessToken),
		xxl.ExecutorPort(conf.ExecutorPort),
		xxl.RegistryKey(conf.RegistryKey),
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
		log.Println(exec.Run())
	}()
}
