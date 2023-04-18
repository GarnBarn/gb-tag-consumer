package gb_tag_consumer

import (
	"fmt"
	"github.com/GarnBarn/common-go/httpserver"
	"github.com/GarnBarn/gb-tag-consumer/config"
	"github.com/sirupsen/logrus"
)

var (
	appConfig config.Config
)

func startHealthCheck() {
	httpServer := httpserver.NewHttpServer()
	logrus.Info("Listening and serving HTTP on :", appConfig.HTTP_SERVER_PORT)
	httpServer.Run(fmt.Sprint(":", appConfig.HTTP_SERVER_PORT))
}

func init() {
	appConfig = config.Load()
}

func main() {
	go startHealthCheck()
}
