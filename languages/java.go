//go:build java
// +build java

package languages

import (
	"fmt"
	"srvexec/common"
)

var (
	MainLanguage = common.Language{
		Name: "python",
		Exec: execute,
	}
)

func execute(j common.ToExecute) (common.Status, string) {
	fmt.Printf("Execute java with %#v\n", j)

	if j.Code == "" {
		return common.ErrorCompile, "No code"
	} else {
		return common.Ok, "Code executed"
	}
}
