package main

// 搜索Class文件
// 加载Hello World之前，首先要加载它的超类，也就是java.lang.Object。
// 在调用main()方法前，因为虚拟机需要准备好参数数组，所以需要加载java.lang.String和
// java.lang.String[]类。把字符串打印到控制台还需要加载java.lang.System类，等等。
// 那么，java虚拟机从哪里寻找这些类呢？

// 先在文件夹中：go build
// 之后，使用该指令可以读取class数据：./ch02 -Xjre "~/JDK/jre" java.lang.Object

import "fmt"
import "strings"
import "../ch02/classpath"

func main() {
	// 解析命令行
	cmd := parseCmd()

	// 如果是版本号
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" { // 帮助命令
		printUsage()
	} else { // 其他
		startJVM(cmd)
	}
}

// 启动JVM虚拟机
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	// 输出class数据
	fmt.Printf("class data:%v\n", classData)
}
