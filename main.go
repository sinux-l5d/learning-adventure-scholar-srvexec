package main

import (
	"srvexec/common"
	"srvexec/environments"
)

func main() {
	app := common.Webserver(environments.MainEnvironments.Handler)
	app.Listen(":8080")
}
