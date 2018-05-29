package main

import "fmt"

func main() {
	// The first go program
	fmt.Println("Hello, World!")
}

/*
从上面简短的代码中，可以看出go包含以下几个部分：
- package声明
- 导入包
- 函数
- 变量
- 表达式
- 注释

main 是一个特殊package的名字，go程序必须包含main包
fmt 包含格式化I/O函数
main（） 主函数，必须以main函数开始，作为执行入口
注释：单行用//来表示，
多行用 斜杠星号+注释内容+星号斜杠组成

Println函数是在Print基础，再向os.Stdout写入一个换行符
它以大写字符开始，表示是被导出的，可以外部访问的函数

go程序不需要用分号结尾。
import两种用法：一种是: import "fmt"
再一种就是：
import(
	"fmt"
	"math/rand"
)

fmt官方文档介绍： https://golang.org/pkg/fmt/
*/
