package judge

import "github.com/si9ma/KillOJ-common/utils"

var (
	CompileSuccessMsg = map[string]string{
		"zh": "编译成功",
		"en": "compile success",
	}

	RunSuccessMsg = map[string]string{
		"zh": "结果正确",
		"en": "Accepted",
	}
)

func GetCompileSuccessMsg() string {
	lang := utils.GetLang()
	if val, ok := CompileSuccessMsg[lang]; ok {
		return val
	}

	return ""
}

func GetRunSuccessMsg() string {
	lang := utils.GetLang()
	if val, ok := RunSuccessMsg[lang]; ok {
		return val
	}

	return ""
}
