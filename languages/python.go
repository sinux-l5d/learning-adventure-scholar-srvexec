//go:build libpython

package languages

import (
	"os/exec"
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

func Execute(code string) (string, error) {
	out, err := exec.Command("python", "-c", code).CombinedOutput()
	return string(out), err
}
