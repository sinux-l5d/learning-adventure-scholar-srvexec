package main

import (
	"srvexec/common"
	"srvexec/environments"
)

func main() {
	app := common.Webserver(environments.MainEnvironments.Handler)
	common.LogFatal(app.Listen("localhost:8080").Error())
}
