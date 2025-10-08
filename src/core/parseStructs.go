package core

import (
	"iter"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/types"
)

type state byte

const (
	NotStarted state = iota
	Started
	Finished
)

// parseStructs parses the tokens of the struct field types.
func (env *Environment) parseStructs(structs iter.Seq[*types.Struct]) error {
	processed := map[*types.Struct]state{}

	for structure := range structs {
		err := env.parseStruct(structure, processed)

		if err != nil {
			return err
		}
	}

	return nil
}

// parseStruct parses the field tokens of a single struct.
func (env *Environment) parseStruct(structure *types.Struct, processed map[*types.Struct]state) error {
	if processed[structure] == Finished {
		return nil
	}

	processed[structure] = Started
	offset := 0

	for i, field := range structure.Fields {
		file := structure.File.(*fs.File)
		typ, err := env.TypeFromTokens(field.Tokens[1:], file)

		if err != nil {
			return err
		}

		dependency, isStruct := typ.(*types.Struct)

		if isStruct {
			switch processed[dependency] {
			case NotStarted:
				err := env.parseStruct(dependency, processed)

				if err != nil {
					return err
				}
			case Started:
				return errors.New(&CycleDetected{A: structure.Name(), B: dependency.Name()}, file, field.Tokens[0].Position)
			case Finished:
			}
		}

		field.Type = typ
		field.Index = uint8(i)
		field.Offset = uint8(offset)
		offset += field.Type.Size()
	}

	processed[structure] = Finished
	return nil
}