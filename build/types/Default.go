package types

// Default represents the default types in our type system.
var Default = map[string]*Type{
	"Byte":    Byte,
	"Int":     Int,
	"Int64":   Int64,
	"Int32":   Int32,
	"Int16":   Int16,
	"Int8":    Int8,
	"Float":   Float,
	"Float64": Float64,
	"Float32": Float32,
	"Pointer": Pointer,
	"Text":    Text,
}
