package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
)

const VERSION = "0.0.1"

var (
	help   bool   // 命令行帮助
	url    string // 下载链接
	thread int    // 下载并发数
)

// init 对程序所需flag信息进行初始化定义
func init() {
	// 定义flag参数
	flag.BoolVar(&help, "h", false, "command line help")
	flag.StringVar(&url, "u", "", "set download `url`, must format as [https/http/ftp]://target.url")
	flag.IntVar(&thread, "t", runtime.NumCPU(), "set downloader `thread`, default is number of logical CPUs")

	// 重写使用帮助参数
	flag.Usage = func() {
		// 打印版本号及用法
		_, _ = fmt.Fprintf(os.Stderr, `Himd version: himd/`+VERSION+`
Usage: himd [-u url] [-t thread]

Options:
`)

		// 打印默认用法
		flag.PrintDefaults()
	}
}

// main 主程序入口，接收命令行参数作为输入
func main() {
	// 解析flag参数
	flag.Parse()

	// 如果解析到了-h参数，打印用法
	if help {
		flag.Usage()
		return
	}

	// 检查输入参数是否符合要求
	verifyFlag()
}

// verifyFlag 检查flag是否符合要求
func verifyFlag() {
	// 如果下载链接未指定，输出错误信息并退出程序
	if url == "" {
		log.Fatalln("error: no specifies download target, `-u` option can not be null")
	}

	// 校验url是否符合要求
	if ret, err := regexp.MatchString("^((https|http|ftp)?://)[^\\s]+", url); !ret {
		if err != nil {
			log.Fatalln("flag verify failed, err: ", err.Error())
		} else {
			log.Fatalln("url must format as [https/http/ftp]://target.url")
		}
	}

	// 校验thread是否符合要求
	if thread <= 0 {
		log.Fatalln("thread num must max than 0")
	}
}
