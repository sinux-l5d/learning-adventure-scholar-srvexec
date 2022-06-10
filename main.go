package main

// Proxy

import (
	"srvexec/common"
	"srvexec/languages"
)

func main() {
	// cherche le language dans la liste
	app := common.Webserver(languages.MainLanguage.Exec)
	app.Listen(":8080")
}
