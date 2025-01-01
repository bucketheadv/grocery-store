package trigger

import (
	"HereWeGo/core/timezone"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
)

func NextTriggerTimes(spec string, t time.Time, loc *time.Location, n int) []time.Time {
	c := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	s, err := c.Parse(spec)
	if err != nil {
		logrus.Error(err)
	}
	dateTime := timezone.WithZone(t, loc)
	var result []time.Time
	for i := 0; i < n; i++ {
		dateTime = s.Next(dateTime)
		result = append(result, dateTime)
	}
	return result
}
