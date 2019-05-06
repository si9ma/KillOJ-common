package judge

import "github.com/si9ma/KillOJ-common/utils"

// 11xx: compile error
const (
	INNER_COMPILER_ERR = 1101 // error from real compiler
	OUTER_COMPILER_ERR = 1102 // error from our outer compiler
	COMPILE_TIMEOUT    = 1103
)

// 1[2-4]xx: container error
const (
	// 12xx: error from outermost process (the process to start container)
	RUNNER_ERR = 1201

	// 13xx: container error (error from container)
	CONTAINER_ERR = 1301

	// 14xx: run error (error from program run in container)
	APP_ERR                   = 1401
	WRONG_ANSWER_ERR          = 1402
	OUT_OF_MEMORY_ERR         = 1403
	RUN_TIMEOUT_ERR           = 1404
	BAD_SYSTEMCALL_ERR        = 1405
	NO_ENOUGH_PID_ERR         = 1406
	JAVA_SECURITY_MANAGER_ERR = 1407
)

type ErrAdapter struct {
	// inner error message
	// lang code -> message
	InnerMsg    map[string]string
	OuterStatus Status // the status expose to user

	// the message expose to user,
	// lang code -> message
	OuterMsg map[string]string
}

// System error
var SystemErrAdapter = ErrAdapter{
	InnerMsg:    SystemErrorMsg,
	OuterStatus: SystemErrorStatus,
	OuterMsg:    SystemErrorMsg,
}

var (
	CompileTimeOutMsg = map[string]string{
		"zh": "编译时间太长，请检查代码!",
		"en": "compile too long, Please check your code!",
	}

	RunTimeOutMsg = map[string]string{
		"zh": "程序运行超时，请检查代码!",
		"en": "run program timeout, Please check your code!",
	}

	SystemErrorMsg = map[string]string{
		"zh": "糟糕，判题机异常，请向管理员报告异常。",
		"en": "Oops, something has gone wrong with the judger. Please report this to administrator.",
	}

	RuntimeErrorMsg = map[string]string{
		"zh": "运行时错误，请检查代码!",
		"en": "Runtime error, Please check your code!",
	}

	CompileErrorMsg = map[string]string{
		"zh": "编译错误,请检查代码!",
		"en": "compile fail,Please check your code!",
	}

	WrongAnswerErrorMsg = map[string]string{
		"zh": "结果错误!",
		"en": "Wrong answer!",
	}

	OOMErrorMsg = map[string]string{
		"zh": "超出内存使用限制!",
		"en": "Memory Limit Exceeded!",
	}

	BadSysErrorMsg = map[string]string{
		"zh": "非法系统调用!",
		"en": "Illegal system call!",
	}

	NoEnoughPidErrorMsg = map[string]string{
		"zh": "超出PID最大允许值限制!",
		"en": "No Enough PID!",
	}

	JavaSecurityManagerErrorMsg = map[string]string{
		"zh": "非法Java操作!",
		"en": "Illegal Java operation!",
	}
)

// map errno to ErrAdapter
var ErrAdapterMapping = map[int64]ErrAdapter{
	RUNNER_ERR:         SystemErrAdapter,
	CONTAINER_ERR:      SystemErrAdapter,
	OUTER_COMPILER_ERR: SystemErrAdapter,

	INNER_COMPILER_ERR: {
		InnerMsg:    CompileErrorMsg,
		OuterStatus: CompileErrorStatus,
		OuterMsg:    CompileErrorMsg,
	},
	COMPILE_TIMEOUT: {
		InnerMsg:    CompileTimeOutMsg,
		OuterStatus: CompileErrorStatus,
		OuterMsg:    CompileTimeOutMsg,
	},
	APP_ERR: {
		InnerMsg:    RuntimeErrorMsg,
		OuterStatus: RuntimeErrorStatus,
		OuterMsg:    RuntimeErrorMsg,
	},
	RUN_TIMEOUT_ERR: {
		InnerMsg:    RunTimeOutMsg,
		OuterStatus: RunTimeOutStatus,
		OuterMsg:    RunTimeOutMsg,
	},
	WRONG_ANSWER_ERR: {
		InnerMsg:    WrongAnswerErrorMsg,
		OuterStatus: WrongAnswerStatus,
		OuterMsg:    WrongAnswerErrorMsg,
	},
	OUT_OF_MEMORY_ERR: {
		InnerMsg:    OOMErrorMsg,
		OuterStatus: OOMStatus,
		OuterMsg:    OOMErrorMsg,
	},

	// InnerMsg and OuterMsg is different
	BAD_SYSTEMCALL_ERR: {
		InnerMsg:    BadSysErrorMsg,
		OuterStatus: RuntimeErrorStatus,
		OuterMsg:    RuntimeErrorMsg,
	},
	NO_ENOUGH_PID_ERR: {
		InnerMsg:    NoEnoughPidErrorMsg,
		OuterStatus: RuntimeErrorStatus,
		OuterMsg:    RuntimeErrorMsg,
	},
	JAVA_SECURITY_MANAGER_ERR: {
		InnerMsg:    JavaSecurityManagerErrorMsg,
		OuterStatus: RuntimeErrorStatus,
		OuterMsg:    RuntimeErrorMsg,
	},
}

// return empty string when don't exist
func GetInnerErrorMsgByErrNo(errno int64) string {
	lang := utils.GetLang()
	if adapter, ok := ErrAdapterMapping[errno]; !ok {
		if val, ok := adapter.InnerMsg[lang]; ok {
			return val
		}
	}

	return ""
}

// return empty string when don't exist
func GetOuterErrorMsgByErrNo(errno int64) string {
	lang := utils.GetLang()
	if adapter, ok := ErrAdapterMapping[errno]; !ok {
		if val, ok := adapter.OuterMsg[lang]; ok {
			return val
		}
	}

	return ""
}

// return empty status when don't exist
func GetStatusByErrNo(errno int64) Status {
	if adapter, ok := ErrAdapterMapping[errno]; !ok {
		return adapter.OuterStatus
	}

	return NullStatus
}
