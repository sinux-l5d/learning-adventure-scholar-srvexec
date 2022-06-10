package common

import "fmt"

func Format(annotation string, multiline string, lang string) {
	fmt.Println(annotation + ":\n\"\"\"" + lang + "\n" + multiline + "\n\"\"\"\n")
}
