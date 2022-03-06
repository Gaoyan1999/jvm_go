package main

import (
	"fmt"
	"jvmgo/ch02/classpath"
	"jvmgo/ch05/classfile"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// // If n < 0, there is no limit on the number of replacements.
	className := strings.Replace(cmd.class, ".", "/", -1)
	classFile := loadClass(className, cp)
	mainMethod := getMainMethod(classFile)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

}
func loadClass(className string, cp *classpath.ClassPath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	if classFile, err := classfile.Parse(classData) ; err == nil {
		return classFile
	}else {
		panic(err)
	}

}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
