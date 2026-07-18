package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// evaluateDot converts a dot expression to an SSA value.
func (f *Function) evaluateDot(expr *expression.Expression) (ssa.Value, error) {
	if len(expr.Children) != 2 {
		return nil, errors.NewAt(MissingFieldName, f.File, expr.Source().End())
	}

	right := expr.Children[1]
	left := expr.Children[0]

	if left.Token.Kind == token.Invalid {
		return nil, errors.NewAt(MissingObject, f.File, left.Source().Start())
	}

	if left.Token.Kind == token.Identifier && left.Token.StringFrom(f.File.Bytes) == "asm" {
		return f.evaluateAsm(right)
	}

	reset := len(f.Block().Instructions)
	leftValue, err := f.evaluate(left)

	if err != nil {
		return nil, err
	}

	pkgValue, isPackage := leftValue.(*ssa.Package)

	if isPackage {
		pkg := f.Env.Package(pkgValue.Name, f.File)

		if pkg == nil {
			return nil, errors.New(&UnknownIdentifier{Name: pkgValue.Name}, f.File, left.Source())
		}

		if right.Token.Kind != token.Identifier {
			return nil, errors.New(ExpectedPackageMember, f.File, right.Source())
		}

		rightText := right.Token.StringFrom(f.File.Bytes)
		return f.evaluatePackageMember(pkg, rightText, expr)
	}

	enumValue, isEnum := leftValue.(*ssa.Enum)

	if isEnum {
		if right.Token.Kind != token.Identifier {
			return nil, errors.New(ExpectedEnumMember, f.File, right.Source())
		}

		rightText := right.Token.StringFrom(f.File.Bytes)
		constExpr, exists := enumValue.Typ.Member(rightText)

		if !exists {
			return nil, errors.New(&UnknownEnumMember{EnumName: enumValue.Typ.Name(), MemberName: rightText}, f.File, right.Source())
		}

		tmp := f.File
		f.File = enumValue.Typ.File().(*fs.File)
		value, err := f.evaluateRight(constExpr.(*expression.Expression))
		f.File = tmp

		if err != nil {
			return nil, err
		}

		return value, nil
	}

	if expr.Parent != nil && expr.Parent.Token.Kind == token.Call && expr.Parent.Children[0] == expr {
		f.Block().Instructions = f.Block().Instructions[:reset]
		return f.evaluateMethod(leftValue, left, right, expr)
	}

	if right.Token.Kind != token.Identifier {
		return nil, errors.New(ExpectedStructField, f.File, right.Source())
	}

	switch leftValue := leftValue.(type) {
	case *ssa.Struct:
		return f.fieldFromStruct(leftValue, left, right)

	case *ssa.Call:
		return f.fieldFromCall(leftValue, left, right, expr)

	default:
		return f.fieldFromMemory(leftValue, left, right, expr)
	}
}