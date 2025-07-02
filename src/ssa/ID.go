package ssa

type Id int

func (id Id) ID() int {
	return int(id)
}

func (id *Id) SetID(newId int) {
	*id = Id(newId)
}