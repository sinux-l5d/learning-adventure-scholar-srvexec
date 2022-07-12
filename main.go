package main

import (
	"srvexec/common"
	"srvexec/environments"
)

func main() {
	app := common.Webserver(environments.MainEnvironments)

	// Allow empty value, app.Listen don't require host
	listen := common.Config.GetDefault("LISTEN", "127.0.0.1")

	// Don't allow empty value, app.Listen require port
	port, set := common.Config.GetSafe("PORT")
	if !set || port == "" {
		port = "8080"
	}

	common.LogInfo("Listening on " + listen + ":" + port)
	common.LogFatal(app.Listen(listen + ":" + port).Error())
}
