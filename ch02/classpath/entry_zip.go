package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry表示ZIP或JAR文件形式的类路径
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) * ZipEntry{
	abs, err := filepath.Abs(path)
	if err!= nil {
		panic(err)
	}
	return &ZipEntry{absPath: abs}
}
func (self * ZipEntry) readClass(className string)([]byte, Entry, error){
	reader, err := zip.OpenReader(self.absPath)
	if err!=nil {
		return nil,nil ,err
	}
	defer reader.Close()
	// jar 包下所有文件
	for _,f := range reader.File {
		// 遍历ZIP压缩包里的文件，看能否找到class文件
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil,nil ,err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil,nil ,err
			}
			return data,self,err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (self *ZipEntry) String() string {
	return self.absPath
}