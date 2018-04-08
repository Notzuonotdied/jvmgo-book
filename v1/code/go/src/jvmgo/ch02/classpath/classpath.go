package classpath

import "os"
import "path/filepath"

// 这里采用了组合模式（composite pattern）来实现。

// 类路径
type Classpath struct {
	// 启动类路径
	bootClasspath Entry
	// 扩展类路径
	extClasspath Entry
	// 用户类路径
	userClasspath Entry
}

// Parse函数使用-Xjre选项解析启动类路径和扩展类路径。
// 使用-classpath/-cp选项解析用户类路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 优先使用用户输入的-Xjre选项作为jre目录。如果没有输入该选项，则在当前目录下
// 选找jre目录。如果找不到，尝试使用JAVA_HOME环境变量。
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

//
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

// 用于判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 解析用户的路径
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

// 如果用户没有提供-classpath/-cp选项，则使用当前目录作为用户类路径。
// ReadClass()方法依次从启动类路径、扩展类路径和用户类路径中搜索class
// 文件。
// className: fully/qualified/ClassName
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
