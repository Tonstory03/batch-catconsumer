package main

import (
	"th.truecorp.it.dsm.batch/batch-catconsumer/apirouter"
	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
	"th.truecorp.it.dsm.batch/batch-catconsumer/cronjob"
)

func main() {

	// loading config.
	config.LoadingConfig()

	// start cronjob
	cronjob.Init()

	//	Start server
	apirouter.SetupAPIRouter()
}
