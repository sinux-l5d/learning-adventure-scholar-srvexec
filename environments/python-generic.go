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

	out, err := languages.Execute(codeFinal)

	if err != nil {
		return string(out), common.ErrorExec
	}

	return string(out), common.Ok
}