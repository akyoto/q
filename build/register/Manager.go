package register

// Manager manages the allocation state of registers.
type Manager struct {
	Registers        []*Register
	SyscallRegisters []*Register
}

// NewManager creates a new register manager.
func NewManager() *Manager {
	manager := &Manager{
		Registers: []*Register{
			{Name: "rbx"},
			{Name: "rbp"},
			{Name: "r12"},
			{Name: "r13"},
			{Name: "r14"},
			{Name: "r15"},

			// These registers are clobbered after syscalls:
			// {Name: "rax"},
			// {Name: "rcx"},
			// {Name: "r11"},
		},
		SyscallRegisters: []*Register{
			{Name: "rax"},
			{Name: "rdi"},
			{Name: "rsi"},
			{Name: "rdx"},
			{Name: "r10"},
			{Name: "r8"},
			{Name: "r9"},
		},
	}

	return manager
}

// FindFreeRegister tries to find a free register
// and returns nil when all are currently occupied.
func (manager *Manager) FindFreeRegister() *Register {
	for _, register := range manager.Registers {
		if register.UsedBy == nil {
			return register
		}
	}

	return nil
}

// RegisterByName returns the register with the given name.
func (manager *Manager) RegisterByName(name string) *Register {
	for _, register := range manager.Registers {
		if register.Name == name {
			return register
		}
	}

	return nil
}
