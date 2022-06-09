package common

type Status uint8

const (
	StatusUndefined Status = iota
	Ok
	Nok
	ErrorCompile
	ErrorExec
)

func (s Status) String() string {
	switch s {
	case Ok:
		return "ok"
	case Nok:
		return "nok"
	case ErrorCompile:
		return "error_compile"
	case ErrorExec:
		return "error_exec"
	default:
		return "undefined"
	}
}
