package cronjob

import (
	"github.com/robfig/cron/v3"
	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
)

func Init() {
	var startJob bool
	c := cron.New()
	cronJobs := config.GetCronJobs()

	for _, cronJob := range cronJobs {

		if !cronJob.Enable {
			continue
		}

		switch cronJob.Name {
		case NAME_RETRY_PROCESS_KAFKA:
			c.AddFunc(cronJob.Expression, retryProcessKafka)
			startJob = true

		}

	}
	if startJob {
		c.Start()
	}
}
