package compiler

import (
	"sort"

	"git.urbach.dev/cli/q/src/core"
)

const (
	// MinimumAssertionDensity is the minimum assertion percentage per package.
	MinimumAssertionDensity = 2

	// MinimumStatements is the minimum number of statements before the density check applies.
	MinimumStatements = 100 / MinimumAssertionDensity
)

// lintAssertionDensity checks that every package has enough assertions.
func lintAssertionDensity(env *core.Environment) error {
	var errs []*LowAssertionDensity

	for name, pkg := range env.Packages {
		asserts := 0
		statements := 0

		for _, function := range pkg.Functions {
			for variant := range function.Variants {
				asserts += int(variant.Count.Assert)
				statements += int(variant.Count.Statement)
			}
		}

		if statements < MinimumStatements {
			continue
		}

		if asserts < statements/MinimumStatements {
			errs = append(errs, &LowAssertionDensity{Package: name, Asserts: asserts, Statements: statements})
		}
	}

	if len(errs) == 0 {
		return nil
	}

	sort.Slice(errs, func(i int, j int) bool {
		return errs[i].Package < errs[j].Package
	})

	return errs[0]
}