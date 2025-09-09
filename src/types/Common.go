package types

var (
	Any        = &Base{name: "any", size: 0}
	AnyInt     = &Base{name: "int"}
	AnyPointer = &Pointer{To: Any}
	Bool       = &Base{name: "bool", size: 1}
	Error      = &Base{name: "error", size: 8}
	Int64      = &Base{name: "int64", size: 8}
	Int32      = &Base{name: "int32", size: 4}
	Int16      = &Base{name: "int16", size: 2}
	Int8       = &Base{name: "int8", size: 1}
	Float64    = &Base{name: "float64", size: 8}
	Float32    = &Base{name: "float32", size: 4}
	Nil        = &Base{name: "nil", size: 8}
	UInt64     = &Base{name: "uint64", size: 8}
	UInt32     = &Base{name: "uint32", size: 4}
	UInt16     = &Base{name: "uint16", size: 2}
	UInt8      = &Base{name: "uint8", size: 1}
	Void       = &Base{name: "void", size: 0}
)

var (
	Byte  = UInt8
	Float = Float64
	Int   = Int64
	UInt  = UInt64
)

var (
	CString = &Pointer{To: Byte}
	String  = &Struct{
		Package:    "",
		UniqueName: "string",
		name:       "string",
		Fields: []*Field{
			{Name: "ptr", Type: CString, Index: 0, Offset: 0},
			{Name: "len", Type: UInt, Index: 1, Offset: 8},
		},
	}
)