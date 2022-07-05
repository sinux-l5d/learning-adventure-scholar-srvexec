package common

type Status uint8

const (
	StatusUndefined Status = iota
	Ok
	Nok
	ErrorCompile
	ErrorExec
	ErrorInternal
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
	case ErrorInternal:
		return "error_internal"
	default:
		return "undefined"
	}
}

func (s Status) HttpCode() int {
	switch s {
	case Ok:
		return 200
	case Nok:
		return 400
	case ErrorCompile:
		return 400
	case ErrorExec:
		return 400
	case ErrorInternal:
		return 500
	default:
		return 500
	}
}

func StatusFromString(s string) Status {
	switch s {
	case "ok":
		return Ok
	case "nok":
		return Nok
	case "error_compile":
		return ErrorCompile
	case "error_exec":
		return ErrorExec
	case "error_internal":
		return ErrorInternal
	default:
		return StatusUndefined
	}
}
