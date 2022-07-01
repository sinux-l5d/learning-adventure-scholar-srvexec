//go:build python_generic

package environments

import (
	"srvexec/common"
	"srvexec/languages"
)

var (
	MainEnvironments = common.Environment{
		Name:    "python-generic",
		Handler: handlePython,
	}
)

type contextPython struct {
	BeforeCode string `json:"beforeCode"`
	AfterCode  string `json:"afterCode"`
	Indent     int    `json:"indent"`
}

func handlePython(j common.ToHandle) (string, common.Status) {
	if j.Code == "" {
		return "No code", common.ErrorCompile
	}

	// Convertion du json en struct
	var ctx contextPython

	if err := j.Exercice.UnmarshalContexte(&ctx); err != nil {
		common.LogError("Error while unmarshalling context: " + err.Error())
		return "Error unmarshalling context", common.ErrorInternal
	}

	// Indendation optionel du code
	codeEtu := languages.Indent(j.Code, ctx.Indent)

	// Incorporation du contexte
	codeFinal := ctx.BeforeCode + codeEtu + ctx.AfterCode

	out, err := languages.Execute(codeFinal)

	if err != nil {
		common.LogError("Error while executing code: " + err.Error())
		return string(out), common.ErrorExec
	}

	return string(out), common.Ok
}
