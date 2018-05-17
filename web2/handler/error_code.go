package handler

var (
	// ErrorCodes contains all error messages and their codes.
	// 400<service code><error code>
	// service code: 2 number
	// error code: 3 number
	ErrorCodes = map[int]string{
		// 账号错误码
		40001000: "用户ID填写错误",
		40001001: "公司名称填写错误",
		40001002: "用户已经存在",
		40001010: "文件格式错误",
		40001011: "用户ID填写错误",
		40001020: "无权限查看",
		40001021: "此用户未上链",

		// 内部错误码
		40002000: "不支持的查询类型",

		// 鉴权错误码
		40003000: "不支持的鉴权方式",
	}
)
