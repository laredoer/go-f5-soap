package utils

import "regexp"

var (
	tempNameRe = regexp.MustCompile(`/(.*)/(.*)`) // 模板名称正则
)

// GetNameByRegexp 通过正则表达式去除 name 前缀
func GetNameByRegexp(name string) string {
	submatch := tempNameRe.FindAllStringSubmatch(name, -1)
	if len(submatch) > 0 && len(submatch[0]) > 1 {
		return submatch[0][2]
	}
	return ""
}
