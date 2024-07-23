package GoCommon_Python

import (
	"fmt"
	"testing"

	"github.com/qaqab/GoCommon_Python"
)

func TestPythonIn(t *testing.T) {
	// 调用 ControlPython 包中的 Python_in 函数，传入两个参数："/root/" 和一个字符串切片，切片中包含 "*.py" 和 "/root/"
	success_type := GoCommon_Python.Python_in("/root/", []string{"*.py", "/root/"})

	// 打印 success_type 的值
	fmt.Println(success_type)
}
func TestPythonSplit(t *testing.T) {
	// 调用ControlPython包中的Python_split函数，传入参数为字符串"/root/asdasd/asdasd/asf"和分隔符"/"，并限制分割后数组的长度为2
	success_list := GoCommon_Python.Python_split("/root/asdasd/asdasd/asf", "/", 2)
	// 打印分割后的数组
	fmt.Println(success_list)
}
