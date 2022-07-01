//go:build libpython

package languages

import (
	"context"
	"os/exec"
	"srvexec/common"
	"strings"
)

// Indente un `code` multiligne d'un nombre `indent` de tabulations
func Indent(code string, indent int) string {
	codeEtu := strings.Split(code, "\n")

	for i, row := range codeEtu {
		codeEtu[i] = strings.Repeat("\t", indent) + row
	}

	return strings.Join(codeEtu, "\n")
}

// Execute le code Python en utilisant python -c.
func Execute(code string) (string, error) {
	// On kill le processus s'il dépasse le timeout
	ctx, cancel := common.GetTimeoutCtx(common.GetTimeoutEnv())
	defer cancel()

	cmd := exec.CommandContext(ctx, "python", "-c", code)

	out, err := cmd.CombinedOutput()

	// Timeout
	if ctx.Err() == context.DeadlineExceeded {
		// Command was killed
		common.LogInfo("Command was killed")
		common.LogDebug("Context error: %v | Command error: %v", ctx.Err(), err)
		return "timeout", ctx.Err()
	}

	// On regarde si l'exécution a échoué
	if err != nil {
		// Command error
		common.LogError("Command error: %v | output: %s", err, common.WrapMultiline(string(out), "output"))
	}

	return string(out), err
}
