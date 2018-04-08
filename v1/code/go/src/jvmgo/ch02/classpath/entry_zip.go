package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	// 存放zip或者jar的绝对路径
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

// 从zip中提取class文件
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 如果打不开zip，就直接返回
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	// 使用defer关键字延迟关闭资源，以免内存泄漏
	// 使用defer语句确保打开的文件得以关闭。
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			// 打开
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			// 关闭
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			// 注意：在循环中不断打开和关闭，效率非常的低，请见entry_zip2.go的优化
			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
