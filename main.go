package main

import (
	"os"
	"srvexec/common"
	"srvexec/environments"
)

func main() {
	app := common.Webserver(environments.MainEnvironments.Handler)

	listen, set := os.LookupEnv("SRVEXEC_LISTEN")
	if !set {
		listen = "127.0.0.1"
	}

	port, set := os.LookupEnv("SRVEXEC_PORT")
	if !set || port == "" {
		port = "8080"
	}

	common.LogFatal(app.Listen(listen + ":" + port).Error())
}
