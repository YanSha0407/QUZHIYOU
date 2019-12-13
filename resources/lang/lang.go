package lang

import (
	"QUZHIYOU/resources/lang/zh_cn"
)

// GetMsg get error information based on Code
func GetMsg(str string) string {
	// make map
	MsgFlags := make(map[string]string)
	// 判断系统配置的语言包类型
	MsgFlags = zh_cn.MsgFlags
	msg, ok := MsgFlags[str]
	if ok {
		return msg
	}
	return str
}
