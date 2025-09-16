package core_test

import (
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/go/assert"
)

var errs = []struct {
	File          string
	ExpectedError error
}{
	{"DefinitionCountMismatch.q", &core.DefinitionCountMismatch{Function: "main.swap", Count: 1, ExpectedCount: 2}},
	{"DefinitionCountMismatch2.q", &core.DefinitionCountMismatch{Function: "main.swap", Count: 3, ExpectedCount: 2}},
	{"DefinitionCountMismatch3.q", &core.DefinitionCountMismatch{Function: "main.nothing", Count: 1, ExpectedCount: 0}},
	{"ErrorNotChecked.q", &core.ErrorNotChecked{Identifier: "value"}},
	{"InvalidCallExpression.q", core.InvalidCallExpression},
	{"InvalidCondition.q", core.InvalidCondition},
	{"InvalidCondition2.q", core.InvalidCondition},
	{"InvalidExpression.q", core.InvalidExpression},
	{"InvalidExpression2.q", core.InvalidExpression},
	{"InvalidExpression3.q", core.InvalidExpression},
	{"InvalidExpression4.q", core.InvalidExpression},
	{"InvalidExpression5.q", core.InvalidExpression},
	{"InvalidExpression6.q", core.InvalidExpression},
	{"InvalidExpression7.q", core.InvalidExpression},
	{"InvalidExpression8.q", core.InvalidExpression},
	{"InvalidFieldInit.q", core.InvalidFieldInit},
	{"InvalidFieldInit2.q", core.InvalidFieldInit},
	{"InvalidLoopHeader.q", core.InvalidLoopHeader},
	{"InvalidStructOperation.q", core.InvalidStructOperation},
	{"InvalidStructOperation2.q", core.InvalidStructOperation},
	{"MissingOperand.q", core.MissingOperand},
	{"NoMatchingFunction.q", &core.NoMatchingFunction{Function: "main.f"}},
	{"NotDataStruct.q", &core.NotDataStruct{TypeName: "int"}},
	{"NotDataStruct2.q", &core.NotDataStruct{TypeName: "int64"}},
	{"ParameterCountMismatch.q", &core.ParameterCountMismatch{Function: "main.f", Count: 0, ExpectedCount: 1}},
	{"ParameterCountMismatch2.q", &core.ParameterCountMismatch{Function: "main.f", Count: 2, ExpectedCount: 1}},
	{"PartiallyUnknownIdentifier.q", &core.PartiallyUnknownIdentifier{Name: "x"}},
	{"ResourceAlreadyConsumed.q", &core.UnknownIdentifier{Name: "x"}},
	{"ResourceAlreadyConsumed2.q", &core.UnknownIdentifier{Name: "x"}},
	{"ResourceAlreadyConsumed3.q", &core.UnknownIdentifier{Name: "x"}},
	{"ResourceAlreadyConsumed4.q", &core.UnknownIdentifier{Name: "x"}},
	{"ResourceNotConsumed.q", &core.ResourceNotConsumed{TypeName: "!int64"}},
	{"ResourcePartiallyConsumed.q", &core.ResourcePartiallyConsumed{TypeName: "!int64"}},
	{"ResourceTypeMismatch.q", &core.TypeMismatch{Encountered: "int64", Expected: "!int64", ParameterName: "_", IsReturn: false}},
	{"ReturnCountMismatch.q", &core.ReturnCountMismatch{Count: 1, ExpectedCount: 0}},
	{"ReturnCountMismatch2.q", &core.ReturnCountMismatch{Count: 1, ExpectedCount: 2}},
	{"ReturnCountMismatch3.q", &core.ReturnCountMismatch{Count: 0, ExpectedCount: 1}},
	{"ReturnCountMismatch4.q", &core.ReturnCountMismatch{Count: 0, ExpectedCount: 1}},
	{"TypeMismatch.q", &core.TypeMismatch{Encountered: "string", Expected: "int64", ParameterName: "x", IsReturn: false}},
	{"TypeMismatch2.q", &core.TypeMismatch{Encountered: "string", Expected: "int64", ParameterName: "y", IsReturn: true}},
	{"TypeMismatch3.q", &core.TypeMismatch{Encountered: "string", Expected: "int64"}},
	{"TypeMismatch4.q", &core.TypeMismatch{Encountered: "string", Expected: "uint8"}},
	{"TypeMismatch5.q", &core.TypeMismatch{Encountered: "!string", Expected: "*any"}},
	{"TypeMismatch6.q", &core.TypeMismatch{Encountered: "int", Expected: "*int64"}},
	{"UndefinedStructField.q", &core.UndefinedStructField{Identifier: "p", FieldName: "y"}},
	{"UnknownIdentifier.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier2.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier3.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier4.q", &core.UnknownIdentifier{Name: "unknown"}},
	{"UnknownIdentifier5.q", &core.UnknownIdentifier{Name: "unknown"}},
	{"UnknownIdentifier6.q", &core.UnknownIdentifier{Name: "run"}},
	{"UnknownIdentifier7.q", &core.UnknownIdentifier{Name: "run.unknown"}},
	{"UnknownIdentifier8.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier9.q", &core.UnknownIdentifier{Name: "err"}},
	{"UnknownIdentifier10.q", &core.UnknownIdentifier{Name: "value"}},
	{"UnknownIdentifier11.q", &core.UnknownIdentifier{Name: "value"}},
	{"UnknownIdentifier12.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownStructField.q", &core.UnknownStructField{StructName: "string", FieldName: "unknown"}},
	{"UnknownType.q", &core.UnknownType{Name: "unknown"}},
	{"UnknownType2.q", &core.UnknownType{Name: "unknown"}},
	{"UnknownType3.q", &core.UnknownType{Name: "unknown"}},
	{"UnknownType4.q", &core.UnknownType{Name: "unknown"}},
	{"UnnecessaryCast.q", core.UnnecessaryCast},
	{"UnusedValue.q", &core.UnusedValue{Value: "42"}},
	{"UnusedValue2.q", &core.UnusedValue{Value: "2 + 3"}},
	{"UnusedValue3.q", &core.UnusedValue{Value: "\"not used\""}},
	{"UnusedValue4.q", &core.UnusedValue{Value: "1"}},
	{"UnusedValue5.q", &core.UnusedValue{Value: "x + 1"}},
	{"UnusedValue6.q", &core.UnusedValue{Value: "x2"}},
	{"VariableAlreadyExists.q", &core.VariableAlreadyExists{Name: "x"}},
	{"VariableAlreadyExists2.q", &core.VariableAlreadyExists{Name: "x"}},
}

func TestErrors(t *testing.T) {
	for _, test := range errs {
		name := strings.TrimSuffix(test.File, ".q")

		t.Run(name, func(t *testing.T) {
			b := config.New(filepath.Join("testdata", test.File))
			_, err := compiler.Compile(b)
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), test.ExpectedError.Error())
		})
	}
}