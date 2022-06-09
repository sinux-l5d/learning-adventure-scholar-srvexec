//go:build python
// +build python

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
	fmt.Printf("Execute python with %#v\n", j)

	if j.Code == "" {
		return common.ErrorCompile, "No code"
	} else {
		return common.Ok, "Code executed"
	}
}
