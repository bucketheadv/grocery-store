package components

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xxl-job/xxl-job-executor-go"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	logrus.Printf("XxlJob日志 - "+format, a...)
}

func (l *logger) Error(format string, a ...interface{}) {
	logrus.Printf("XxlJob日志 - "+format, a...)
}

var exec xxl.Executor

func init() {
	conf := GetConfig().XxlJob
	// 开始日志清理
	go startClearLogFile()

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
		var dateTime = time.UnixMilli(req.LogDateTim)
		filePath := fmt.Sprintf("%s%c%s", conf.LogDir, os.PathSeparator, dateTime.Format(time.DateOnly))
		if _, err := os.Stat(filePath); err != nil {
			logrus.Errorf("读取目录异常, %s", err.Error())
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				logrus.Printf("尝试创建目录失败, %s", err.Error())
			}
		}

		var fileName = fmt.Sprintf("%s%c%d.log", filePath, os.PathSeparator, req.LogID)
		content, err := readLogFile(fileName, req.FromLineNum)
		if err != nil {
			logrus.Errorf("读取文件异常: %s", err.Error())
		}
		return &xxl.LogRes{
			Code: 200,
			Msg:  "测试消息",
			Content: xxl.LogResContent{
				FromLineNum: req.FromLineNum,
				ToLineNum:   2,
				LogContent:  content,
				IsEnd:       true,
			},
		}
	})

	go func() {
		logrus.Error(exec.Run())
	}()
}

func LogJobInfo(param *xxl.RunReq, format string, a ...interface{}) {
	conf := GetConfig().XxlJob
	dir := fmt.Sprintf("%s%c%s", conf.LogDir, os.PathSeparator, time.Now().Format(time.DateOnly))
	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			logrus.Errorf("创建文件失败, %s", err.Error())
			return
		}
	}
	fileName := fmt.Sprintf("%s%c%d.log", dir, os.PathSeparator, param.LogID)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm); err != nil {
			logrus.Errorf("创建文件失败, %s", err.Error())
			return
		}
	}
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		logrus.Errorf("打开文件失败, %s", err.Error())
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logrus.Error(err.Error())
		}
	}(f)
	prefix := fmt.Sprintf("[%s] ", time.Now().Format(time.DateTime))
	logData := prefix + fmt.Sprintf(format, a...) + "\n"
	_, _ = f.WriteString(logData)
}

func readLogFile(filePath string, fromLineNo int) (string, error) {
	if filePath == "" {
		return "", errors.New("文件名错误")
	}
	if _, err := os.Stat(filePath); err != nil {
		return "", err
	}
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logrus.Error(err.Error())
		}
	}(f)
	scanner := bufio.NewScanner(f)
	lineNum := 0
	lines := make([]string, 0)
	for scanner.Scan() {
		lineNum++
		if lineNum >= fromLineNo {
			lines = append(lines, scanner.Text())
		}
	}
	return strings.Join(lines, "\n"), nil
}

func RegTask(pattern string, taskFunc xxl.TaskFunc) {
	if exec != nil {
		exec.RegTask(pattern, taskFunc)
	}
}

func startClearLogFile() {
	tick := time.Tick(1 * time.Hour)
	for {
		select {
		case <-tick:
			logrus.Info("开始清理Xxl-Job日志")
			conf := GetConfig().XxlJob
			logRetention := conf.LogRetention
			whiteListLogs := make([]string, 0)
			for i := 0; i <= logRetention; i++ {
				date := time.Now().Add(-time.Duration(i) * 24 * time.Hour).Format(time.DateOnly)
				whiteListLogs = append(whiteListLogs, date)
			}
			err := filepath.Walk(conf.LogDir, func(path string, info os.FileInfo, err error) error {
				if path == conf.LogDir || info == nil {
					return nil
				}
				fullPath := strings.Replace(path, conf.LogDir, conf.LogDir, -1)
				if info.IsDir() && !slices.Contains(whiteListLogs, info.Name()) {
					if err := os.RemoveAll(fullPath); err != nil {
						logrus.Error(err.Error())
					}
				}
				return nil
			})
			if err != nil {
				logrus.Error(err)
			}
			break
		}
	}
}
