package build

import (
	"sort"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// Scope represents a map of variables.
type Scope map[string]*Variable

// ScopeStack represents a list of scopes.
type ScopeStack struct {
	scopes []Scope
}

// ScopeError represents a scope error.
type ScopeError struct {
	Err      error
	Position token.Position
}

// Add adds a variable.
func (stack *ScopeStack) Add(variable *Variable) {
	top := stack.scopes[len(stack.scopes)-1]
	top[variable.Name] = variable
}

// Get returns the variable with the specified name.
func (stack *ScopeStack) Get(variableName string) *Variable {
	for index := len(stack.scopes) - 1; index >= 0; index-- {
		scope := stack.scopes[index]
		variable := scope[variableName]

		if variable != nil {
			return variable
		}
	}

	return nil
}

// Each executes a function on each variable.
func (stack *ScopeStack) Each(callBack func(*Variable)) {
	for index := len(stack.scopes) - 1; index >= 0; index-- {
		scope := stack.scopes[index]

		for _, variable := range scope {
			callBack(variable)
		}
	}
}

// Push pushes a new scope to the top of the stack.
func (stack *ScopeStack) Push() {
	stack.scopes = append(stack.scopes, Scope{})
}

// Pop removes the scope at the top of the stack.
func (stack *ScopeStack) Pop() {
	stack.scopes = stack.scopes[:len(stack.scopes)-1]
}

// Errors returns a list of errors at the top of the stack.
func (stack *ScopeStack) Errors(isLoop bool) []*ScopeError {
	var scopeErrors []*ScopeError
	scope := stack.scopes[len(stack.scopes)-1]

	for _, variable := range scope {
		if !variable.Used {
			scopeErrors = append(scopeErrors, &ScopeError{
				Position: variable.Position,
				Err:      &errors.UnusedVariable{VariableName: variable.Name},
			})
		}

		if !variable.LastAssignUsed {
			scopeErrors = append(scopeErrors, &ScopeError{
				Position: variable.LastAssign,
				Err:      &errors.IneffectiveAssignment{VariableName: variable.Name},
			})
		}

		if variable.Mutable && variable.LastAssign == variable.Position {
			scopeErrors = append(scopeErrors, &ScopeError{
				Position: variable.Position,
				Err:      &errors.UnmodifiedMutable{VariableName: variable.Name},
			})
		}
	}

	if len(scopeErrors) == 0 {
		return nil
	}

	sort.Slice(scopeErrors, func(a int, b int) bool {
		return scopeErrors[a].Position < scopeErrors[b].Position
	})

	return scopeErrors
}
