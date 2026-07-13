package linter_test

import (
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/linter"
	"git.urbach.dev/go/assert"
)

var errs = []struct {
	File          string
	ExpectedError error
}{
	{"AlwaysTrue.q", linter.AlwaysTrue},
	{"AlwaysTrue2.q", linter.AlwaysTrue},
	{"AlwaysTrue3.q", linter.AlwaysTrue},
	{"AlwaysTrue4.q", linter.AlwaysTrue},
	{"AlwaysFalse.q", linter.AlwaysFalse},
	{"AlwaysFalse2.q", linter.AlwaysFalse},
	{"AlwaysFalse3.q", linter.AlwaysFalse},
	{"AlwaysFalse4.q", linter.AlwaysFalse},
	{"IdenticalExpressions.q", &linter.IdenticalExpressions{Operator: "-"}},
	{"IdenticalExpressions2.q", &linter.IdenticalExpressions{Operator: "/"}},
	{"IdenticalExpressions3.q", &linter.IdenticalExpressions{Operator: "%"}},
	{"IdenticalExpressions4.q", &linter.IdenticalExpressions{Operator: "&"}},
	{"IdenticalExpressions5.q", &linter.IdenticalExpressions{Operator: "|"}},
	{"IdenticalExpressions6.q", &linter.IdenticalExpressions{Operator: "^"}},
	{"IdenticalExpressions7.q", &linter.IdenticalExpressions{Operator: "=="}},
	{"IdenticalExpressions8.q", &linter.IdenticalExpressions{Operator: "!="}},
	{"IdenticalExpressions9.q", &linter.IdenticalExpressions{Operator: "<"}},
	{"IdenticalExpressions10.q", &linter.IdenticalExpressions{Operator: "<="}},
	{"IdenticalExpressions11.q", &linter.IdenticalExpressions{Operator: ">"}},
	{"IdenticalExpressions12.q", &linter.IdenticalExpressions{Operator: ">="}},
	{"MixedSignedUnsigned.q", &linter.MixedSignedUnsigned{Signed: "int64", Unsigned: "uint64"}},
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