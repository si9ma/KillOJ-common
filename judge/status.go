package judge

type Status string

const (
	RuntimeErrorStatus = Status("RuntimeError")
	CompileErrorStatus = Status("CompileError")
	JudingStatus       = Status("Juding")
	AcceptedStatus     = Status("Accepted")
	RunTimeOutStatus   = Status("RunTimeOut")
	OOMStatus          = Status("OOM")
	WrongAnswerStatus  = Status("WrongAnswer")
	SystemErrorStatus  = Status("SystemError")
)
