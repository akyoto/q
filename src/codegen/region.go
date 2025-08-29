package codegen

// region is used to mark the start and end indices of a block.
type region struct {
	Start uint32
	End   uint32
}