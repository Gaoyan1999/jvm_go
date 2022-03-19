package heap

import (
	"fmt"
	"jvmgo/ch08/classfile"
	"jvmgo/ch08/classpath"
)

type ClassLoader struct {
	classPath   *classpath.ClassPath
	verboseFlag bool
	classMap    map[string]*Class // loaded classes
}

func NewClassLoader(classPath *classpath.ClassPath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		classPath:   classPath,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
}

func (classLoader *ClassLoader) LoadClass(name string) *Class {
	if class, ok := classLoader.classMap[name]; ok {
		return class
	}
	if name[0] == '[' {
		return classLoader.loadArrayClass(name)
	}
	return classLoader.loadNonArrayClass(name)
}

// 加载数组类
func (classLoader *ClassLoader) loadNonArrayClass(name string) *Class {
	// 读取class数据到内存中
	data, entry := classLoader.readClass(name)
	// 解析class文件，生成虚拟机可用到类数据，放入方法区中
	class := classLoader.defineClass(data)
	// 链接
	link(class)
	if classLoader.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

func (classLoader *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		AccessFlags: ACC_PUBLIC, //TODO
		Name:        name,
		InitStarted: true,
		superClass:  classLoader.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			classLoader.LoadClass("java/lang/Cloneable"),
			classLoader.LoadClass("java/io/Serializable"),
		},
	}
	classLoader.classMap[name] = class
	return class
}

func (classLoader *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	class, entry, err := classLoader.classPath.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException:" + name)
	}
	return class, entry
}

func (classLoader *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = classLoader
	resolveSuperClass(class)
	resolveInterfaces(class)
	classLoader.classMap[class.Name] = class
	return class
}

func parseClass(data []byte) *Class {
	classFile, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(classFile)
}

func resolveSuperClass(class *Class) {
	if class.Name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify()
	prepare(class)
}
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func verify() {
	// todo
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
		}
		if field.isLongOrDouble() {
			slotId++
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticFieldSlots = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	staticSlots := class.staticFieldSlots
	constantPool := class.constantPool
	cpIndex := field.constValueIndex
	slotId := field.slotId
	if cpIndex > 0 {
		switch field.descriptor {
		case "Z", "B", "C", "S", "I":
			val := constantPool.GetConstant(cpIndex).(int32)
			staticSlots.SetInt(slotId, val)
		case "J":
			val := constantPool.GetConstant(cpIndex).(int64)
			staticSlots.SetLong(slotId, val)
		case "F":
			val := constantPool.GetConstant(cpIndex).(float32)
			staticSlots.SetFloat(slotId, val)
		case "D":
			val := constantPool.GetConstant(cpIndex).(float64)
			staticSlots.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := constantPool.GetConstant(cpIndex).(string)
			jString := JString(class.loader, goStr)
			staticSlots.SetRef(slotId,jString)
		default:
			panic("TODO")
		}

	}

}
