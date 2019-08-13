package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

const VERSION = "0.0.1"

var (
	help   bool
	url    string
	thread int
)

func init() {
	flag.BoolVar(&help, "h", false, "command line help")
	flag.StringVar(&url, "u", "", "set download `url`, must format as [http/https]://target.url")
	flag.IntVar(&thread, "t", runtime.NumCPU(), "set downloader `thread`, default is number of logical CPUs")

	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, `himd version: himd/`+VERSION+`
Usage: himd [-u url] [-t thread]

Options:
`)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
	}
}
