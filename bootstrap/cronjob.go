package bootstrap

import (
	"strings"

	"github.com/robfig/cron/v3"
)

// RegisterCronjob register cronjob
func RegisterCronjob(c *cron.Cron, t string, f func()) {
	cronjob := strings.Split(t, ",")
	for _, ct := range cronjob {
		c.AddFunc(ct, f)
	}
}
