package common

// Une fonction de type Executor prend en paramètre un dictionnaire (json) et renvoie un status et des logs (optionels).
type Executor func(j ToExecute) (s Status, log string)

type Language struct {
	Name string
	Exec Executor
}

type ToExecute struct {
	Code    string `json:"code"`
	Context string `json:"context"`
}
