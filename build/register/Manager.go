package register

// Manager manages the allocation state of registers.
type Manager struct {
	All         List
	General     List
	Call        List
	Syscall     List
	ReturnValue List
}

// NewManager creates a new register manager.
func NewManager() *Manager {
	// Rather than doing lots of mini allocations
	// we'll allocate memory for all registers at once.
	registers := []Register{
		{ID: 0, Name: "rax"},
		{ID: 1, Name: "rbx"},
		{ID: 2, Name: "rcx"},
		{ID: 3, Name: "rdx"},
		{ID: 4, Name: "rdi"},
		{ID: 5, Name: "rsi"},
		{ID: 6, Name: "rbp"},
		{ID: 7, Name: "r8"},
		{ID: 8, Name: "r9"},
		{ID: 9, Name: "r10"},
		{ID: 10, Name: "r11"},
		{ID: 11, Name: "r12"},
		{ID: 12, Name: "r13"},
		{ID: 13, Name: "r14"},
		{ID: 14, Name: "r15"},
	}

	// To simplify the lists below,
	// bind the registers to their name.
	rax := &registers[0]
	rbx := &registers[1]
	rcx := &registers[2]
	rdx := &registers[3]
	rdi := &registers[4]
	rsi := &registers[5]
	rbp := &registers[6]
	r8 := &registers[7]
	r9 := &registers[8]
	r10 := &registers[9]
	r11 := &registers[10]
	r12 := &registers[11]
	r13 := &registers[12]
	r14 := &registers[13]
	r15 := &registers[14]

	// Register configuration
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
			rdi,
			rsi,
			rdx,
			r10,
			r8,
			r9,
		},
		Syscall: List{
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

// ByID returns the register with the given ID.
func (manager *Manager) ByID(id byte) *Register {
	return manager.All[id]
}
