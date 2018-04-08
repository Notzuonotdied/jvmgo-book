package classpath

import "io/ioutil"
import "path/filepath"

// Go没有专门的构造函数，本书统一使用new开头的函数类创建结构体实例。

// 表示目录形式的路径
type DirEntry struct {
	// 存放绝对路径
	absDir string
}

// 使用本函数创建不同类型的Entry类型
func newDirEntry(path string) *DirEntry {
	// 绝对路径和错误信息，说明：Go语言支持同时返回多个变量
	// newDirEntry先把参数转换为绝对路径，如果转换过程中出现错误，
	// 就调用panic函数终止程序执行。
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	// 返回实例
	return &DirEntry{absDir}
}

// readClass先把目录和class文件名拼成一个完整的路径，
// 然后调用ioutil包提供的ReadFile函数读取class文件的内容
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

// 直接返回目录（绝对路径）
func (self *DirEntry) String() string {
	return self.absDir
}
