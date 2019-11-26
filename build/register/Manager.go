package register

// Manager manages the allocation state of registers.
type Manager struct {
	General     []*Register
	Call        []*Register
	Syscall     []*Register
	ReturnValue []*Register
}

// NewManager creates a new register manager.
func NewManager() *Manager {
	rax := &Register{Name: "rax"}
	rbx := &Register{Name: "rbx"}
	rcx := &Register{Name: "rcx"}
	rdx := &Register{Name: "rdx"}
	rdi := &Register{Name: "rdi"}
	rsi := &Register{Name: "rsi"}
	rbp := &Register{Name: "rbp"}
	r8 := &Register{Name: "r8"}
	r9 := &Register{Name: "r9"}
	r10 := &Register{Name: "r10"}
	r11 := &Register{Name: "r11"}
	r12 := &Register{Name: "r12"}
	r13 := &Register{Name: "r13"}
	r14 := &Register{Name: "r14"}
	r15 := &Register{Name: "r15"}

	manager := &Manager{
		General: []*Register{
			rbx,
			rbp,
			r12,
			r13,
			r14,
			r15,
		},
		Syscall: []*Register{
			rax,
			rdi,
			rsi,
			rdx,
			r10,
			r8,
			r9,
		},
		ReturnValue: []*Register{
			rax,
			rcx,
			r11,
		},
	}

	manager.Call = manager.Syscall
	return manager
}

// FindFreeRegister tries to find a free register
// and returns nil when all are currently occupied.
func (manager *Manager) FindFreeRegister() *Register {
	for _, register := range manager.General {
		if register.IsFree() {
			return register
		}
	}

	return nil
}

// RegisterByName returns the register with the given name.
func (manager *Manager) RegisterByName(name string) *Register {
	for _, register := range manager.General {
		if register.Name == name {
			return register
		}
	}

	return nil
}
