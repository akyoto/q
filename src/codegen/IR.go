package codegen

import "git.urbach.dev/cli/q/src/ssa"

// IR is the intermediate representation for the codegen phase.
type IR struct {
	Steps         []*Step
	ValueToStep   map[ssa.Value]*Step
	BlockToRegion map[*ssa.Block]region
}