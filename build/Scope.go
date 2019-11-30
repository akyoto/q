package build

import "sort"

// Scope represents a map of variables.
type Scope map[string]*Variable

// ScopeStack represents a list of scopes.
type ScopeStack struct {
	scopes []Scope
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

// Unused returns a list of unused variables at the top of the stack.
func (stack *ScopeStack) Unused() []*Variable {
	var unused []*Variable
	scope := stack.scopes[len(stack.scopes)-1]

	for _, variable := range scope {
		if !variable.Used() {
			unused = append(unused, variable)
		}
	}

	if len(unused) == 0 {
		return nil
	}

	sort.Slice(unused, func(a int, b int) bool {
		return unused[a].Position < unused[b].Position
	})

	return unused
}
