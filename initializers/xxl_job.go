package initializers

import (
	"github.com/sirupsen/logrus"
	"github.com/xxl-job/xxl-job-executor-go"
)

type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	logrus.Printf("自定义日志 - "+format, a...)
}

func (l *logger) Error(format string, a ...interface{}) {
	logrus.Printf("自定义日志 - "+format, a...)
}

var exec xxl.Executor

func init() {
	conf := GetConfig().XxlJob
	if !conf.Enabled {
		return
	}

	exec = xxl.NewExecutor(
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
			Msg:  "测试消息",
			Content: xxl.LogResContent{
				FromLineNum: req.FromLineNum,
				ToLineNum:   2,
				LogContent:  "这是自定义日志Handler",
				IsEnd:       true,
			},
		}
	})

	go func() {
		logrus.Error(exec.Run())
	}()
}

func RegTask(pattern string, taskFunc xxl.TaskFunc) {
	if exec != nil {
		exec.RegTask(pattern, taskFunc)
	}
}
