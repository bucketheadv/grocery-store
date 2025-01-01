package trigger

import (
	"HereWeGo/core/timezone"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
)

// NextTriggerTimes 计算cron下次执行时间
// startTime 计算从哪个时间开始后的时间
// loc cron对应计算的时区
// n 返回最近的多少条时间
func NextTriggerTimes(spec string, startTime time.Time, loc *time.Location, n int) []time.Time {
	c := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	s, err := c.Parse(spec)
	if err != nil {
		logrus.Error(err)
	}
	dateTime := timezone.WithZone(startTime, loc)
	var result []time.Time
	for i := 0; i < n; i++ {
		dateTime = s.Next(dateTime)
		result = append(result, dateTime)
	}
	return result
}
