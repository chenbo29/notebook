package common

import (
	"fmt"
	"runtime"
	"strings"
)

// GetPackagePath 获取当前代码所在的包路径
func GetPackagePath() string {
	_, filename, _, _ := runtime.Caller(1)
	parts := strings.Split(filename, "/")
	if len(parts) > 0 {
		// 假设项目路径不包含标准库路径
		return strings.Join(parts[:len(parts)-1], "/")
	}
	return ""
}

// GetPackageName 获取当前代码所在的包名
func GetPackageName() string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	parts := strings.Split(function.Name(), "/")
	if len(parts) > 0 {
		fullName := parts[len(parts)-1]
		packageParts := strings.Split(fullName, ".")
		if len(packageParts) > 1 {
			return packageParts[0]
		}
	}
	return ""
}

func PrintDescription(text any) {
	fmt.Println(strings.Repeat(" ", 4), "【详情】", text)
}
func PrintParams(text any) {
	fmt.Println(strings.Repeat(" ", 4), "【参数】", text)
}

func PrintResult(text any) {
	fmt.Println(strings.Repeat(" ", 4), "【结果】", text, strings.Repeat("\n", 2))
}
