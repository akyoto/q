package build

type Register struct {
	Name   string
	UsedBy *Variable
}

type RegisterManager struct {
	Registers        []*Register
	SyscallRegisters []*Register
}

func NewRegisterManager() *RegisterManager {
	manager := &RegisterManager{
		Registers: []*Register{
			{Name: "rcx"},
			{Name: "rbx"},
			{Name: "rbp"},
			{Name: "r11"},
			{Name: "r12"},
			{Name: "r13"},
			{Name: "r14"},
			{Name: "r15"},
		},
	}

	manager.SyscallRegisters = []*Register{
		{Name: "rax"},
		{Name: "rdi"},
		{Name: "rsi"},
		{Name: "rdx"},
		{Name: "r10"},
		{Name: "r8"},
		{Name: "r9"},
	}

	return manager
}

// FindFreeRegister tries to find a free register
// and returns nil when all are currently occupied.
func (manager *RegisterManager) FindFreeRegister() *Register {
	for _, register := range manager.Registers {
		if register.UsedBy == nil {
			return register
		}
	}

	return nil
}

// RegisterByName returns the register with the given name.
func (manager *RegisterManager) RegisterByName(name string) *Register {
	for _, register := range manager.Registers {
		if register.Name == name {
			return register
		}
	}

	return nil
}
