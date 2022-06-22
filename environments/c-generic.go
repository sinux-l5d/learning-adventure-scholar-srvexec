//go:build c_generic

package environments

import (
	"srvexec/common"
	"srvexec/languages"
)

var (
	MainEnvironments = common.Environment{
		Name:    "c-generic",
		Handler: handleC,
	}
)

type contextC struct {
	Libs []string `json:"libs"`
}

func handleC(j common.ToHandle) (string, common.Status) {
	if j.Code == "" {
		return "No code", common.ErrorCompile
	}

	// Convertion du json en struct
	var ctx contextC

	if j.Exercice.UnmarshalContexte(&ctx) != nil {
		return "Error unmarshalling context", common.ErrorInternal
	}

	// Ajoute les libs au code
	codeFinal := languages.FormatLibs(ctx.Libs) + j.Code

	out, err := languages.Execute(codeFinal)

	if err != common.Ok {
		return out, err
	}

	return out, common.Ok

}
