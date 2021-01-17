package config

/**
 * @Author: yirufeng
 * @Date: 2021/1/16 4:17 下午
 * @Desc:
 **/

var (
	//通用状态码
	SUCCESS = 1000
	FAIL    = 5000

	//用户模块的错误状态码
	//用户名已存在
	UserNameExist            = 5100
	UserPasswordHashedFailed = 5101
	UserNameNotExist         = 5102
	UserPasswordNotRight     = 5103
	UserParamFormatInvalid   = 5104 //传入的参数格式不正确，例如传id但是里面有不能转换为整形的字符串

	//分类模块的状态码

	//文章模块的状态码

	//JWT的token的错误状态码
	TokenFormatInvalid = 5400 //token格式错误
	TokenInValid       = 5401 //token非法
)

var msg = map[int]string{
	1000: "操作成功!",
	5000: "操作失败, 请稍后再试!",

	5100: "用户名已存在!",
	5101: "用户密码加密失败, 请稍后再试!",
	5102: "用户名不存在!",
	5103: "密码不正确!",
}

func GetMessage(code int) string {
	return msg[code]
}
