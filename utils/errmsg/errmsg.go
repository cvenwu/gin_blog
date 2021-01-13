package errmsg

/**
 * @Author: yirufeng
 * @Date: 2021/1/13 3:43 下午
 * @Desc:
 **/

//声明一系列常量
const (
	SUCCESS = 200
	ERROR   = 500

	//区分一下状态码：因为有3个模块，用户，文章，分类
	//	code = 1000... 表示用户模块的错误

	ERROR_USERNAME_USED    = 1001 //用户名已经被使用了
	ERROR_PASSWORD_WRONG   = 1002 //密码错误
	ERROR_USER_NOT_EXIST   = 1003 //用户不存在
	ERROR_TOKEN_EXIST      = 1004 //用户在JWT验证的时候token不存在
	ERROR_TOKEN_RUNTIME    = 1005 //用户在JWT验证的时候token超时
	ERROR_TOKEN_WRONG      = 1006 //用户在JWT验证的时候token不一样
	ERROR_TOKEN_TYPE_WRONG = 1007 //TOKEN格式不正确

	//	code = 2000... 表示文章模块的错误
	//  code = 3000... 表示分类模块的错误

)

//字典：对应的状态码以及要抛出的错误信息
var codemsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在!",
	ERROR_PASSWORD_WRONG:   "密码错误!",
	ERROR_USER_NOT_EXIST:   "用户不存在!",
	ERROR_TOKEN_EXIST:      "Token不存在!",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期!",
	ERROR_TOKEN_WRONG:      "TOKEN不正确!",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

//输出错误信息的函数
func GetErrMsg(code int) string {
	return codemsg[code]
}
