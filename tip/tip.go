package tip

import (
	"github.com/si9ma/KillOJ-common/lang"
	"golang.org/x/text/language"
)

type Tip map[string]string

var (
	// for judger and sandbox
	CompileTimeOutTip = Tip{
		language.Chinese.String(): "编译时间太长，请检查代码!",
		language.English.String(): "compile too long, Please check your code!",
	}

	RunTimeOutTip = Tip{
		language.Chinese.String(): "程序运行超时，请检查代码!",
		language.English.String(): "run program timeout, Please check your code!",
	}

	SystemErrorTip = Tip{
		language.Chinese.String(): "糟糕，判题机异常，请向管理员报告异常。",
		language.English.String(): "Oops, something has gone wrong with the judger. Please report this to administrator.",
	}

	RuntimeErrorTip = Tip{
		language.Chinese.String(): "运行时错误，请检查代码!",
		language.English.String(): "Runtime error, Please check your code!",
	}

	CompileErrorTip = Tip{
		language.Chinese.String(): "编译错误,请检查代码!",
		language.English.String(): "compile fail,Please check your code!",
	}

	WrongAnswerErrorTip = Tip{
		language.Chinese.String(): "结果错误!",
		language.English.String(): "Wrong answer!",
	}

	OOMErrorTip = Tip{
		language.Chinese.String(): "超出内存使用限制!",
		language.English.String(): "Memory Limit Exceeded!",
	}

	BadSysErrorTip = Tip{
		language.Chinese.String(): "非法系统调用!",
		language.English.String(): "Illegal system call!",
	}

	NoEnoughPidErrorTip = Tip{
		language.Chinese.String(): "超出PID最大允许值限制!",
		language.English.String(): "No Enough PID!",
	}

	JavaSecurityManagerErrorTip = Tip{
		language.Chinese.String(): "非法Java操作!",
		language.English.String(): "Illegal Java operation!",
	}

	CompileSuccessTip = Tip{
		language.Chinese.String(): "编译成功",
		language.English.String(): "compile success",
	}

	RunSuccessTip = Tip{
		language.Chinese.String(): "结果正确",
		language.English.String(): "Accepted",
	}

	// for backend
	ArgValidateFailTip = Tip{
		language.Chinese.String(): "参数验证失败",
		language.English.String(): "validate argument fail",
	}

	BadRequestGeneralTip = Tip{
		language.Chinese.String(): "bad request",
		language.English.String(): "bad request",
	}

	ValidateRequireTip = Tip{
		language.Chinese.String(): "%s不能为空",
		language.English.String(): "%s is required",
	}

	ValidateMaxTip = Tip{
		language.Chinese.String(): "%s不能大于%s",
		language.English.String(): "%s cannot be longer than %s",
	}

	ValidateMinTip = Tip{
		language.Chinese.String(): "%s必须大于%s",
		language.English.String(): "%s must be longer than %s",
	}

	ValidateEmailTip = Tip{
		language.Chinese.String(): "非法邮箱地址",
		language.English.String(): "Invalid email format",
	}

	ValidateLenTip = Tip{
		language.Chinese.String(): "%s长度必须为%s",
		language.English.String(): "%s must be %s characters long",
	}

	ValidateInvalidTip = Tip{
		language.Chinese.String(): "%s非法参数",
		language.English.String(): "%s is not valid",
	}

	InternalServerErrorTip = Tip{
		language.Chinese.String(): "系统错误，请联系管理员",
		language.English.String(): "Internal server error, Please contact administrator",
	}

	UserAlreadyExistInOrgTip = Tip{
		language.Chinese.String(): "用户 %s 已经在 %s 中",
		language.English.String(): "user %s already exist in %s",
	}

	OrgShouldExistWhenNoExistTip = Tip{
		language.Chinese.String(): "当填写组织内ID时组织名不能为空",
		language.English.String(): "organization name should not empty when no_in_organization is not empty",
	}

	AlreadyExistTip = Tip{
		language.Chinese.String(): "%s 已经存在",
		language.English.String(): "%s already exist",
	}

	NotExistTip = Tip{
		language.Chinese.String(): "%s 不存在",
		language.English.String(): "%s not exist",
	}

	UserNotExistTip = Tip{
		language.Chinese.String(): "用户 %s 不存在",
		language.English.String(): "user %s not exist",
	}

	PasswordWrong = Tip{
		language.Chinese.String(): "密码错误",
		language.English.String(): "password is wrong",
	}

	ThirdAuthFailTip = Tip{
		language.Chinese.String(): "%s 认证失败",
		language.English.String(): "%s unauthorized fail",
	}

	NoSignupTip = Tip{
		language.Chinese.String(): "你还未注册",
		language.English.String(): "haven't signup",
	}

	NotSupportProviderTip = Tip{
		language.Chinese.String(): "不支持 %s 登录",
		language.English.String(): "provider %s is not supported",
	}

	UnauthorizedGeneralTip = Tip{
		language.Chinese.String(): "unauthorized",
		language.English.String(): "unauthorized",
	}
)

func (t Tip) String() string {
	lan := lang.GetLangFromEnv()
	if val, ok := t[lan.String()]; ok && val != "" {
		return val
	}

	// if not found ,return en version
	return t[language.English.String()]
}

func (t Tip) String2Lang(lang language.Tag) string {
	if val, ok := t[lang.String()]; ok && val != "" {
		return val
	}

	// if not found ,return english version
	return t[language.English.String()]
}
