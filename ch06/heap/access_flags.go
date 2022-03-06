package heap

const (
	ACC_PUBLIC       = 0x0001 // class field method
	ACC_PRIVATE      = 0x0002 //       field method
	ACC_PROTECTED    = 0x0004 //       field method
	ACC_STATIC       = 0x0008 //       field method
	ACC_FINAL        = 0x0010 // class field method
	ACC_SUPER        = 0x0020 // class
	ACC_SYNCHRONIZED = 0x0020 //             method
	ACC_VOLATILE     = 0x0040 //       field
	ACC_BRIDGE       = 0x0040 //             method
	ACC_TRANSIENT    = 0x0080 //       field
	ACC_VARARGS      = 0x0080 //             method
	ACC_NATIVE       = 0x0100 //             method
	ACC_INTERFACE    = 0x0200 // class
	ACC_ABSTRACT     = 0x0400 // class       method
	ACC_STRICT       = 0x0800 //             method
	STRICT           = 0x0800 //             method
	ACC_SYNTHETIC    = 0x1000 // class field method
	ACC_ANNOTATION   = 0x2000 // class
	ACC_ENUM         = 0x4000 // class field
)

type AccessFlags uint16

func (flags AccessFlags) IsPublic() bool       { return flags&ACC_PUBLIC != 0 }
func (flags AccessFlags) IsPrivate() bool      { return flags&ACC_PRIVATE != 0 }
func (flags AccessFlags) IsProtected() bool    { return flags&ACC_PROTECTED != 0 }
func (flags AccessFlags) IsStatic() bool       { return flags&ACC_STATIC != 0 }
func (flags AccessFlags) IsFinal() bool        { return flags&ACC_FINAL != 0 }
func (flags AccessFlags) IsSuper() bool        { return flags&ACC_SUPER != 0 }
func (flags AccessFlags) IsSynchronized() bool { return flags&ACC_SYNCHRONIZED != 0 }
func (flags AccessFlags) IsVolatile() bool     { return flags&ACC_VOLATILE != 0 }
func (flags AccessFlags) IsBridge() bool       { return flags&ACC_BRIDGE != 0 }
func (flags AccessFlags) IsTransient() bool    { return flags&ACC_TRANSIENT != 0 }
func (flags AccessFlags) IsVarargs() bool      { return flags&ACC_VARARGS != 0 }
func (flags AccessFlags) IsNative() bool       { return flags&ACC_NATIVE != 0 }
func (flags AccessFlags) IsInterface() bool    { return flags&ACC_INTERFACE != 0 }
func (flags AccessFlags) IsAbstract() bool     { return flags&ACC_ABSTRACT != 0 }
func (flags AccessFlags) IsStrict() bool       { return flags&ACC_STRICT != 0 }
func (flags AccessFlags) Strict() bool         { return flags&STRICT != 0 }
func (flags AccessFlags) IsSynthetic() bool    { return flags&ACC_SYNTHETIC != 0 }
func (flags AccessFlags) IsAnnotation() bool   { return flags&ACC_ANNOTATION != 0 }
func (flags AccessFlags) IsEnum() bool         { return flags&ACC_ENUM != 0 }
