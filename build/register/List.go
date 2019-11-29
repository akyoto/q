package register

// List is a list of registers.
type List []*Register

// FindFree tries to find a free register
// and returns nil when all are currently occupied.
func (registers List) FindFree() *Register {
	for _, register := range registers {
		if register.IsFree() {
			return register
		}
	}

	return nil
}

// InUse returns a list of registers that are currently in use.
func (registers List) InUse() []*Register {
	var inUse []*Register

	for _, register := range registers {
		if !register.IsFree() {
			inUse = append(inUse, register)
		}
	}

	return inUse
}

// ByName returns the register with the given name.
func (registers List) ByName(name string) *Register {
	for _, register := range registers {
		if register.Name == name {
			return register
		}
	}

	return nil
}
