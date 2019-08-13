package main

import (
	"log"
	"regexp"
)

// verifyFlag 检查flag是否符合要求
func verifyFlag(url string, _ int) {
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
