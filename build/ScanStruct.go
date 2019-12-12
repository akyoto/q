package build

import (
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// scanStruct scans a data structure.
func (file *File) scanStruct(tokens token.List, index token.Position) (*Struct, token.Position, error) {
	var (
		blockLevel = 0
		structure  = &Struct{}
		field      *Field
	)

	index++
	name := tokens[index]

	if name.Kind != token.Identifier {
		return structure, index, NewError(errors.MissingStructName, file.path, tokens[:index+1], nil)
	}

	structure.Name = name.Text()
	index++

	for ; index < len(tokens); index++ {
		t := tokens[index]

		switch t.Kind {
		case token.Identifier:
			if field == nil {
				field = &Field{
					Name: t.Text(),
				}

				continue
			}

			if field.Type == nil {
				typeName := t.Text()
				field.Type = file.environment.Types[typeName]

				if field.Type == nil {
					return structure, index, NewError(&errors.UnknownType{Name: typeName}, file.path, tokens[:index], nil)
				}

				continue
			}

		case token.NewLine:
			if field == nil {
				continue
			}

			if field.Type == nil {
				return structure, index, NewError(&errors.MissingType{Of: field.Name}, file.path, tokens[:index], nil)
			}

			structure.Fields = append(structure.Fields, field)
			structure.Size += field.Type.Size
			field = nil

		case token.BlockStart:
			blockLevel++

		case token.BlockEnd:
			blockLevel--

			if blockLevel != 0 {
				continue
			}

			return structure, index, nil
		}
	}

	return nil, index, nil
}
