package biz_err

const (
	Success     = 200
	UnKnowErr   = 00000
	ServerError = 10000
	BadRequest  = 20000

	SendImgVerificationCodeFail   = 10001
	SendEmailVerificationCodeFail = 10002

	PluginLoadFailed        = 11000 // 插件加载失败
	PluginInitFailed        = 11001 // 插件初始化失败
	PluginUnloadFailed      = 11002 // 插件卸载失败
	PluginNotFound          = 11003 // 插件未找到
	PluginAlreadyExists     = 11004 // 插件已存在
	PluginDependencyError   = 11005 // 插件依赖错误
	PluginMarketError       = 11006 // 插件市场错误
	PluginDownloadError     = 11007 // 插件下载错误
	PluginVerificationError = 11008 // 插件校验错误
	PluginRepoNotFound      = 11009 // 仓库未找到
	PluginRepoExists        = 11010 // 仓库已存在
	PluginChecksumError     = 11011 // 校验和不匹配
)

var CodeMsg = map[int]string{
	Success:     "请求成功",
	UnKnowErr:   "未知业务异常",
	ServerError: "服务端异常",
	BadRequest:  "错误请求",

	SendImgVerificationCodeFail:   "发送图形验证码失败",
	SendEmailVerificationCodeFail: "发送邮箱验证码失败",
}

func GetMessage(code int) string {
	if msg, ok := CodeMsg[code]; ok {
		return msg
	}
	return CodeMsg[UnKnowErr]
}
