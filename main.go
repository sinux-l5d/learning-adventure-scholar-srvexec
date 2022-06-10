package main

// Proxy

import (
	"srvexec/common"
	"srvexec/languages"
)

func main() {
	app := common.Webserver(languages.MainLanguage.Exec)
	app.Listen(":8080")
}
