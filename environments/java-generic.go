//go:build java_generic

package environments

import (
	"fmt"
	"srvexec/common"
)

var (
	MainEnvironments = common.Environment{
		Name: "java",
		Exec: executeJava,
	}
)

func executeJava(j common.ToExecute) (common.Status, string) {
	fmt.Printf("Execute java with %#v\n", j)

	if j.Code == "" {
		return common.ErrorCompile, "No code"
	} else {
		return common.Ok, "Code executed"
	}
}
