package main

import "flag"
import "fmt"
import "os"

// java [-options] class [args...]
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	// Java虚拟机将使用JDK的启动类了路径来寻找和加载Java标准库中的类，因此需要
	// 某种方式制定jre目录的位置。命令行是个不错的选择，所以增加一个非标准的选项
	// -Xjre
	XjreOption string
	class      string
	args       []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	// 解析Jre的路径
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
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}
