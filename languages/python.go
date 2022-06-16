//go:build libpython

package languages

import "strings"

// Indente un `code` multiligne d'un nombre `indent` de tabulations
func Indent(code string, indent int) string {
	codeEtu := strings.Split(code, "\n")

	for i, row := range codeEtu {
		codeEtu[i] = strings.Repeat("\t", indent) + row
	}

	return strings.Join(codeEtu, "\n")
}
