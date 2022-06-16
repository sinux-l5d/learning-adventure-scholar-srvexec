//go:build python_generic

package environments

import (
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

	if j.Exercice.UnmarshalContexte(&ctx) != nil {
		return common.ErrorInternal, "Error unmarshalling context"
	}

	// Indendation optionel du code
	codeEtu := languages.Indent(j.Code, ctx.Indent)

	// Incorporation du contexte

	codeFinal := ctx.BeforeCode + codeEtu + ctx.AfterCode

	out, err := exec.Command("python", "-c", codeFinal).CombinedOutput()

	if err != nil {
		return common.ErrorExec, string(out)
	}

	return common.Ok, string(out)
}
