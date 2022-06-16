package main

import (
	"srvexec/common"
	"srvexec/environments"
)

func main() {
	app := common.Webserver(environments.MainEnvironments.Exec)
	app.Listen(":8080")
}
