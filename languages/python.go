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
		common.Format("CODE", j.Context+"\n"+j.Code, "python")
		out, err := exec.Command("python", "-c", j.Context+"\n"+j.Code).CombinedOutput()

		if err != nil {
			common.Format("OUTPUT", string(out), "raw")
			return common.ErrorExec, string(out)
		}

		return common.Ok, string(out)
	}
}
