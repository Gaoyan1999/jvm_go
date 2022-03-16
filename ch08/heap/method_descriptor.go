package heap

type TypeDescriptor string

type MethodDescriptor struct {
	parameterTypes []TypeDescriptor
	returnType TypeDescriptor
}
