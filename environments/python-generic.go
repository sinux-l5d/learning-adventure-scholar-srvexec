//go:build python_generic

package environments

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"srvexec/common"
	"srvexec/languages"
)

var (
	MainEnvironments = common.Environment{
		Name: "python-generic",
		Exec: executePython,
	}
)

type contextPython struct {
	BeforeCode string `json:"beforeCode"`
	AfterCode  string `json:"afterCode"`
	Indent     int    `json:"indent"`
}

func executePython(j common.ToExecute) (common.Status, string) {
	if j.Code == "" {
		return common.ErrorCompile, "No code"
	}

	// Convertion du json en struct
	var ctx contextPython

	if err := json.Unmarshal(j.Exercice.Contexte, &ctx); err != nil {
		return common.ErrorInternal, fmt.Sprintf("Error unmarshalling context: %s", err)
	}

	// Indendation optionel du code
	codeEtu := languages.Indent(j.Code, ctx.Indent)

	// Incorporation du contexte

	codeFinal := ctx.BeforeCode + codeEtu + ctx.AfterCode

	out, err := exec.Command("python", "-c", codeFinal).CombinedOutput()

	common.Format("INPUT", codeFinal, "py")
	common.Format("OUTPUT", string(out), "raw")
	if err != nil {
		return common.ErrorExec, string(out)
	}
	return common.Ok, string(out)
}
