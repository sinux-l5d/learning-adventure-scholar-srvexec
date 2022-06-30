//go:build libc

package languages

import (
	"os"
	"os/exec"
	"srvexec/common"
	"strings"
)

// Exécute le code C en utilisant le compilateur gcc.
// Le code est passé via stdin, et mis dans un fichier temporaire.
// Le fichier temporaire est nommé avec un hash du code initiale,
// et est supprimé après l'exécution.
func Execute(code string, flags []string) (string, common.Status) {
	// Hash du code, pour le nom du fichier
	h := common.Hash(code)
	prefix := func(msg string) string {
		return "[" + h + "] " + msg
	}

	common.LogInfo(prefix("Received code : " + common.WrapMultiline(code, "c")))

	// Make args
	args := append([]string{"-x", "c", "-", "-o", h + ".out"}, flags...)
	// Configure gcc
	gcc := exec.Command("gcc", args...)
	defer os.Remove(h + ".out")
	defer common.LogInfo(prefix("Removed file"))

	// On passe le code via stdin
	gcc.Stdin = strings.NewReader(code)

	// On lance la compilation
	common.LogInfo(prefix("Compiling code"))
	out, err := gcc.CombinedOutput()

	// On regarde si la compilation a échoué
	if err != nil {
		common.LogError(prefix("Compilation error: " + common.WrapMultiline(string(out)+"\n"+err.Error(), "output")))
		return string(out), common.ErrorCompile
	}

	// On lance l'exécution du binnaire créé
	common.LogInfo(prefix("Executing"))
	out, err = exec.Command("./" + h + ".out").CombinedOutput()

	// On regarde si l'exécution a échoué
	if err != nil {
		common.LogError(prefix("Execution error: " + common.WrapMultiline(string(out), "output")))
		return string(out), common.ErrorExec
	}

	common.LogInfo(prefix("Execution success: " + common.WrapMultiline(string(out), "output")))
	return string(out), common.Ok
}

// Format un tableau de nom de librairie en une chaine de caractères multiligne.
// Ajoute des #include <...>.
func FormatLibs(libs []string) string {
	var libsStr string
	for _, lib := range libs {
		libsStr += "#include <" + lib + ".h>\n"
	}
	return libsStr
}
