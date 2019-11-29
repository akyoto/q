package register

// Manager manages the allocation state of registers.
type Manager struct {
	All         List
	General     List
	Call        List
	ReturnValue List
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
		All: List{
			rax,
			rbx,
			rcx,
			rdx,
			rdi,
			rsi,
			rbp,
			r8,
			r9,
			r10,
			r11,
			r12,
			r13,
			r14,
			r15,
		},
		General: List{
			rbx,
			rbp,
			r12,
			r13,
			r14,
			r15,
		},
		Call: List{
			rax,
			rdi,
			rsi,
			rdx,
			r10,
			r8,
			r9,
		},
		ReturnValue: List{
			rax,
			rcx,
			r11,
		},
	}

	return manager
}
