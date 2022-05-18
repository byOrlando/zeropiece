package douyin

import (
	"fmt"
	"github.com/dop251/goja"
	"io/ioutil"
	"os"
)

func _A() string {
	file, err := os.Open("./XO/JSFile/_signture.js")
	if err != nil {
		if err != nil {
			fmt.Println("打开文件运行出错")
		}
	}
	fd, err := ioutil.ReadAll(file)
	vm := goja.New()
	_, err = vm.RunString(string(fd))
	if err != nil {
		fmt.Println("运行出错vm")
	}
	var fn func(int, string) string
	err = vm.ExportTo(vm.Get("get_sign"), &fn)
	if err != nil {
		fmt.Println("运行出错 fn")
	}
	defer file.Close()
	return fn(123, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")

}
