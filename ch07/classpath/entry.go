package classpath

import (
	"os"
	"strings"
)

// 获取当前os的分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 寻找和加载class文件
	readClass(className string) ([]byte, Entry, error)
	// String（）方法的作用相当于Java中的toString（）
	String() string
}

// 根据 path 解析不同的 entry
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		// java -cp classes;lib\* ...
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
