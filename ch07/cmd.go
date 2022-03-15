package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag         bool
	versionFlag      bool
	verboseClassFlag bool
	verboseInstFlag  bool
	cpOption         string
	//我们的Java虚拟机将使用JDK的启动类路径来寻找和加载Java标准库中的类，因此需要某种方式指定jre目录的位置
	XjreOption string
	class      string
	args       []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	// flag 可以帮助我们处理命令行选项。
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.verboseInstFlag, "log", false, "print instruction log")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}
func printUsage() {
	// os包定义了一个Args变量，其中存放传递给命令行的全部参数
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
