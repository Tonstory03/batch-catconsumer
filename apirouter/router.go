package apirouter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
	"th.truecorp.it.dsm.batch/batch-catconsumer/utils"
)

func SetupAPIRouter() {

	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	port := config.GetServer().Port

	if utils.IsEmptyString(&port) {
		// setting default port
		port = "8080"
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}
