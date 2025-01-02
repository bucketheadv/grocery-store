package initializer

import (
	"HereWeGo/conf"
	"github.com/bucketheadv/infragin/components"
)

var XxlJobClient components.XxlJobClient

func init() {
	config := conf.Config.XxlJob
	client := components.NewJobClient(config)
	XxlJobClient = client
}
