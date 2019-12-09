package types

var (
	Int64 = &Type{Name: "Int64", Size: 8}
	Int32 = &Type{Name: "Int32", Size: 4}
	Int16 = &Type{Name: "Int16", Size: 2}
	Int8  = &Type{Name: "Int8", Size: 1}
	Byte  = Int8
	Int   = Int64
)
