package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

func main() {

	print(`系统类型：`)
	println(runtime.GOOS)

	print(`系统架构：`)
	println(runtime.GOARCH)

	print(`CPU 核数：`)
	println(runtime.GOMAXPROCS(0))

	print(`电脑名称：`)
	name, err := os.Hostname()
	if err == nil {
		println(name)
	} else {
		println(err)
	}
}

/**
 * 获取电脑CPUId
 */
func getCpuId() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	str := string(out)
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	return str[11:]
}
