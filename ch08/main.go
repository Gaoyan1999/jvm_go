package main

import (
	"fmt"
	"jvmgo/ch08/classpath"
	"jvmgo/ch08/heap"
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
	loader := heap.NewClassLoader(cp,cmd.verboseInstFlag)
	className := strings.Replace(cmd.class,".","/",-1)
	mainClass := loader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod !=nil {
		interpret(mainMethod,cmd.verboseInstFlag,cmd.args)
	}else {
		fmt.Printf("Main method not found in class %s\n",cmd.class)
	}
}