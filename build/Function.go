package build

import (
	"fmt"
	"sync"

	"github.com/akyoto/q/build/assembler"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
	"github.com/akyoto/q/build/types"
)

// Function represents a function.
type Function struct {
	Name             string
	Parameters       []*Parameter
	ReturnTypes      []*types.Type
	ReturnTypeTokens []token.Token
	File             *File
	TokenStart       token.Position
	TokenEnd         token.Position
	Error            error
	NoParameterCheck bool
	IsBuiltin        bool
	IsFinished       bool
	SideEffects      int32
	CallCount        int32
	Finished         *sync.Cond
	FinishedMutex    sync.Mutex
	assembler        *assembler.Assembler
	parameterStart   token.Position
	returnTypeStart  token.Position
}

// Tokens returns all tokens within the function body (excluding the braces '{' and '}').
func (function *Function) Tokens() []token.Token {
	return function.File.tokens[function.TokenStart:function.TokenEnd]
}

// NewError creates an error inside the function.
func (function *Function) NewError(position token.Position, err error) error {
	metaError, hasMetaData := err.(*Error)

	if hasMetaData {
		return metaError
	}

	return NewError(err, function.File.path, function.File.tokens[:function.TokenStart+position+1], function)
}

// CanInline returns true if the function call can be inlined.
func (function *Function) CanInline() bool {
	return len(function.assembler.Instructions) <= 4
}

// InlineInto adds the assembler instructions to another function.
// It excludes the starting label and the last return statement.
func (function *Function) InlineInto(other *Function) {
	// NOTE: We should re-set the register pointers for these instructions
	// because the assembly optimizer relies on pointer equality checks.
	inlinedInstructions := function.assembler.Instructions[1 : len(function.assembler.Instructions)-1]
	other.assembler.Instructions = append(other.assembler.Instructions, inlinedInstructions...)
}

// HasReturnValue returns true if the function has a return value.
func (function *Function) HasReturnValue() bool {
	return len(function.ReturnTypes) > 0
}

// NewErrorf creates a formatted error inside the function.
func (function *Function) NewErrorf(position token.Position, message string, args ...interface{}) error {
	return function.NewError(position, fmt.Errorf(message, args...))
}

// UsedRegisterIDs returns the IDs of used registers.
func (function *Function) UsedRegisterIDs() []register.ID {
	if function.IsBuiltin {
		// return map[string]struct{}{
		// 	// Parameters
		// 	"rax": {},
		// 	"rdi": {},
		// 	"rsi": {},
		// 	"rdx": {},
		// 	"r10": {},
		// 	"r8":  {},
		// 	"r9":  {},

		// 	// Clobbered registers
		// 	"rcx": {},
		// 	"r11": {},
		// }
		return nil
	}

	return function.assembler.UsedRegisterIDs()
}

// Wait will block until the compilation finishes.
func (function *Function) Wait() {
	if function.IsBuiltin {
		return
	}

	function.Finished.L.Lock()
	defer function.Finished.L.Unlock()

	for !function.IsFinished {
		function.Finished.Wait()
	}
}

// String returns the function name.
func (function *Function) String() string {
	return function.Name
}
