package classpath

import "os"
import "strings"

// :(linux/unix) or ;(windows)
// 常量pathListSeparator是string类型，存放路径分割符，文件名有.class后缀。
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// className: fully/qualified/ClassName.class
	readClass(className string) ([]byte, Entry, error)
	String() string
}

// 创建不同类型的Entry
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		//
		return newZipEntry(path)
	}
	// 目录形式的类路径
	return newDirEntry(path)
}
