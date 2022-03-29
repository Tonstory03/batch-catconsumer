package main

import (
	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
	"th.truecorp.it.dsm.batch/batch-catconsumer/cronjob"
)

func main() {

	// loading config.
	config.LoadingConfig()

	// //	Start server
	// apirouter.SetupAPIRouter()

	// start cronjob
	cronjob.Init()
}
