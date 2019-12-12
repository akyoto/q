package build

import (
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
	"github.com/akyoto/q/build/types"
)

// scanStruct scans a data structure.
func (file *File) scanStruct(tokens token.List, index token.Position) (*types.Type, token.Position, error) {
	var (
		blockLevel = 0
		typ        = &types.Type{}
		field      *types.Field
	)

	index++
	name := tokens[index]

	if name.Kind != token.Identifier {
		return typ, index, NewError(errors.MissingStructName, file.path, tokens[:index+1], nil)
	}

	typ.Name = name.Text()
	index++

	for ; index < len(tokens); index++ {
		t := tokens[index]

		switch t.Kind {
		case token.Identifier:
			if field == nil {
				field = &types.Field{
					Name: t.Text(),
				}

				continue
			}

			if field.Type == nil {
				typeName := t.Text()
				field.Type = file.environment.Types[typeName]

				if field.Type == nil {
					return typ, index, NewError(&errors.UnknownType{Name: typeName}, file.path, tokens[:index], nil)
				}

				continue
			}

		case token.NewLine:
			if field == nil {
				continue
			}

			if field.Type == nil {
				return typ, index, NewError(&errors.MissingType{Of: field.Name}, file.path, tokens[:index], nil)
			}

			typ.Fields = append(typ.Fields, field)
			typ.Size += field.Type.Size
			field = nil

		case token.BlockStart:
			blockLevel++

		case token.BlockEnd:
			blockLevel--

			if blockLevel != 0 {
				continue
			}

			return typ, index, nil
		}
	}

	return nil, index, nil
}
