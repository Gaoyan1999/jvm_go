package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath Entry
	userClassPath Entry
}

// 使用 -Xjre 选项解析启动类路径和扩展类路径
func Parse(jreOption,cpOption string) * ClassPath{
	cp := &ClassPath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 如果用户没有提供-classpath/-cp选项，则使用当前目录作为用户类路径。
//ReadClass（）方法依次从启动类路径、扩展类路径和用户类路径中搜索class文件
func (self * ClassPath) ReadClass(className string)([]byte,Entry,error){
	className += ".class"
	//java/lang/Object.class
	if data, entry, err := self.bootClassPath.readClass(className); err == nil {
		return data,entry,nil
	}
	if data, entry, err := self.extClassPath.readClass(className); err == nil {
		return data,entry,nil
	}
	return self.userClassPath.readClass(className)

}

func (self * ClassPath) String() string {
 return self.userClassPath.String()
}

func (self * ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath =  newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	extPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClassPath =  newWildcardEntry(extPath)
}

func (self* ClassPath) parseUserClasspath(option string) {
 if option == "" {
 	option = "."
 }
 self.userClassPath = newEntry(option)
}

func getJreDir(jreOption string) string {
	// 优先使用用户输入的-Xjre选项作为jre目录。如果没有输入该选项，则在当前目录下寻找jre目录
	if jreOption !="" && exists(jreOption){
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	//如果找不到，尝试使用JAVA_HOME环境变量。
	if jh := os.Getenv("JAVA_HOME"); jh !="" {
		return filepath.Join(jh,"jre")
	}
	panic("Can not find jre folder!")
}

// 判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err!=nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
