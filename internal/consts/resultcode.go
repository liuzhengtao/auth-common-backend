package consts

const (
	SUCCESS                          = "00000"
	USER_ERROR                       = "A0001"
	REPEAT_SUBMIT_ERROR              = "A0002"
	USER_LOGIN_ERROR                 = "A0200"
	USER_NOT_EXIST                   = "A0201"
	USER_ACCOUNT_LOCKED              = "A0202"
	USER_ACCOUNT_INVALID             = "A0203"
	USER_PASSWORD_ERROR              = "A0204"
	USERNAME_OR_PASSWORD_ERROR       = "A0210"
	PASSWORD_ENTER_EXCEED_LIMIT      = "A0211"
	CLIENT_AUTHENTICATION_FAILED     = "A0212"
	VERIFY_CODE_TIMEOUT              = "A0213"
	VERIFY_CODE_ERROR                = "A0214"
	TOKEN_INVALID                    = "A0230"
	TOKEN_ACCESS_FORBIDDEN           = "A0231"
	AUTHORIZED_ERROR                 = "A0300"
	ACCESS_UNAUTHORIZED              = "A0301"
	FORBIDDEN_OPERATION              = "A0302"
	PARAM_ERROR                      = "A0400"
	RESOURCE_NOT_FOUND               = "A0401"
	PARAM_TRANSFORM_ERROR            = "A0402"
	PARAM_IS_NULL                    = "A0410"
	USER_UPLOAD_FILE_ERROR           = "A0700"
	USER_UPLOAD_FILE_TYPE_NOT_MATCH  = "A0701"
	USER_UPLOAD_FILE_SIZE_EXCEEDS    = "A0702"
	USER_UPLOAD_IMAGE_SIZE_EXCEEDS   = "A0703"
	SYSTEM_EXECUTION_ERROR           = "B0001"
	SYSTEM_EXECUTION_TIMEOUT         = "B0100"
	SYSTEM_ORDER_PROCESSING_TIMEOUT  = "B0101"
	SYSTEM_DISASTER_RECOVERY_TRIGGER = "B0200"
	FLOW_LIMITING                    = "B0210"
	DEGRADATION                      = "B0220"
	SYSTEM_RESOURCE_ERROR            = "B0300"
	SYSTEM_RESOURCE_EXHAUSTION       = "B0310"
	SYSTEM_RESOURCE_ACCESS_ERROR     = "B0320"
	SYSTEM_READ_DISK_FILE_ERROR      = "B0321"
	CALL_THIRD_PARTY_SERVICE_ERROR   = "C0001"
	MIDDLEWARE_SERVICE_ERROR         = "C0100"
	INTERFACE_NOT_EXIST              = "C0113"
	MESSAGE_SERVICE_ERROR            = "C0120"
	MESSAGE_DELIVERY_ERROR           = "C0121"
	MESSAGE_CONSUMPTION_ERROR        = "C0122"
	MESSAGE_SUBSCRIPTION_ERROR       = "C0123"
	MESSAGE_GROUP_NOT_FOUND          = "C0124"
	DATABASE_ERROR                   = "C0300"
	DATABASE_TABLE_NOT_EXIST         = "C0311"
	DATABASE_COLUMN_NOT_EXIST        = "C0312"
	DATABASE_DUPLICATE_COLUMN_NAME   = "C0321"
	DATABASE_DEADLOCK                = "C0331"
	DATABASE_PRIMARY_KEY_CONFLICT    = "C0341"
)

var ResultCode = map[string]string{
	SUCCESS:                          "成功",
	USER_ERROR:                       "用户端错误",
	REPEAT_SUBMIT_ERROR:              "您的请求已提交，请不要重复提交或等待片刻再尝试。",
	USER_LOGIN_ERROR:                 "用户登录异常",
	USER_NOT_EXIST:                   "用户不存在",
	USER_ACCOUNT_LOCKED:              "用户账户被冻结",
	USER_ACCOUNT_INVALID:             "用户账户已作废",
	USER_PASSWORD_ERROR:              "密码错误",
	PASSWORD_ENTER_EXCEED_LIMIT:      "密码输入次数超限，请稍后再试。",
	CLIENT_AUTHENTICATION_FAILED:     "客户端认证失败",
	VERIFY_CODE_TIMEOUT:              "验证码已过期，请重新获取。",
	VERIFY_CODE_ERROR:                "验证码错误，请重新输入。",
	TOKEN_INVALID:                    "token无效或已过期",
	TOKEN_ACCESS_FORBIDDEN:           "token已被禁止访问",
	AUTHORIZED_ERROR:                 "访问权限异常",
	ACCESS_UNAUTHORIZED:              "访问未授权",
	FORBIDDEN_OPERATION:              "禁止操作",
	PARAM_ERROR:                      "用户请求参数错误",
	PARAM_TRANSFORM_ERROR:            "参数转换异常",
	RESOURCE_NOT_FOUND:               "请求资源不存在",
	PARAM_IS_NULL:                    "请求必填参数为空",
	USER_UPLOAD_FILE_ERROR:           "用户上传文件异常",
	USER_UPLOAD_FILE_TYPE_NOT_MATCH:  "用户上传文件类型不匹配",
	USER_UPLOAD_FILE_SIZE_EXCEEDS:    "用户上传文件太大",
	USER_UPLOAD_IMAGE_SIZE_EXCEEDS:   "用户上传图片太大",
	SYSTEM_EXECUTION_ERROR:           "系统执行出错",
	SYSTEM_EXECUTION_TIMEOUT:         "系统执行超时",
	SYSTEM_ORDER_PROCESSING_TIMEOUT:  "系统订单处理超时",
	SYSTEM_DISASTER_RECOVERY_TRIGGER: "系统容灾功能被出发",
	FLOW_LIMITING:                    "系统限流",
	DEGRADATION:                      "系统功能降级",
	SYSTEM_RESOURCE_ERROR:            "系统资源异常",
	SYSTEM_RESOURCE_EXHAUSTION:       "系统资源耗尽",
	SYSTEM_RESOURCE_ACCESS_ERROR:     "系统资源访问异常",
	SYSTEM_READ_DISK_FILE_ERROR:      "读取磁盘文件失败",
	CALL_THIRD_PARTY_SERVICE_ERROR:   "调用第三方服务异常",
	MIDDLEWARE_SERVICE_ERROR:         "中间件服务异常",
	INTERFACE_NOT_EXIST:              "接口不存在",
	MESSAGE_SERVICE_ERROR:            "消息服务异常",
	MESSAGE_DELIVERY_ERROR:           "消息投递异常",
	MESSAGE_CONSUMPTION_ERROR:        "消息消费异常",
	MESSAGE_SUBSCRIPTION_ERROR:       "消息订阅异常",
	MESSAGE_GROUP_NOT_FOUND:          "消息分组未查到",
	DATABASE_ERROR:                   "数据库服务出错",
	DATABASE_TABLE_NOT_EXIST:         "数据表不存在",
	DATABASE_COLUMN_NOT_EXIST:        "数据列不存在",
	DATABASE_DUPLICATE_COLUMN_NAME:   "多表关联中存在多个相同名称的列",
	DATABASE_DEADLOCK:                "数据库死锁",
	DATABASE_PRIMARY_KEY_CONFLICT:    "数据主键冲突",
}

func GetResultMessage(code string) string {
	if v, ok := ResultCode[code]; ok {
		return v
	}
	return ""
}
