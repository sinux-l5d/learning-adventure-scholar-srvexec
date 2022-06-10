//go:build python

package languages

import (
	"fmt"
	"os/exec"
	"srvexec/common"
)

var (
	MainLanguage = common.Language{
		Name: "python",
		Exec: executePython,
	}
)

func executePython(j common.ToExecute) (common.Status, string) {
	fmt.Printf("Execute python with %#v\n", j)

	if j.Code == "" {
		return common.ErrorCompile, "No code"
	} else {
		out, err := exec.Command("python", "-c", j.Code).Output()

		if err != nil {
			fmt.Printf("Exec Error: %s\n", err)
			return common.ErrorExec, err.Error()
		}

		return common.Ok, string(out)
	}
}
