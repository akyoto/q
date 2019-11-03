package main

import (
	"fmt"
	"sort"

	"github.com/akyoto/q/similarity"
	"github.com/akyoto/q/spec"
	"github.com/akyoto/q/token"
)

// handleToken processes a single token.
func (compiler *File) handleToken(t token.Token) error {
	fmt.Println(t.Kind, string(t.Text))

	switch t.Kind {
	// "Hello"
	case token.Text:
		// text := string(token.Text[1 : len(token.Text)-1])
		// compiler.assembler.Println(text)

	case token.StartOfLine:
		if len(compiler.funcCalls) > 0 {
			return compiler.Error("Missing closing bracket ')'")
		}

	// '('
	case token.ParenthesesStart:
		previous := compiler.PreviousToken(-1)

		if previous.Kind != token.Identifier {
			return compiler.Error("Expected function name before '('")
		}

		isDefinition := false

		if len(compiler.Tokens) >= 3 {
			whiteSpace := compiler.PreviousToken(-2)
			functionKeyword := compiler.PreviousToken(-3)

			if whiteSpace.Kind == token.WhiteSpace && functionKeyword.Kind == token.Keyword && string(functionKeyword.Text) == "func" {
				isDefinition = true
			}
		}

		if isDefinition {

		} else {
			compiler.funcCalls = append(compiler.funcCalls, FunctionCall{
				FunctionName:   string(previous.Text),
				ParameterStart: len(compiler.Tokens) + 1,
			})
		}

	// ')'
	case token.ParenthesesEnd:
		if len(compiler.funcCalls) == 0 {
			return compiler.Error("Missing opening bracket '('")
		}

		call := compiler.funcCalls[len(compiler.funcCalls)-1]

		// Add the last parameter
		if call.ParameterStart < len(compiler.Tokens) {
			call.Parameters = append(call.Parameters, compiler.Tokens[call.ParameterStart:])
		}

		// Currently, we only allow builtin functions
		if !spec.Functions[call.FunctionName] {
			knownFunctions := []string{"print"}

			// Suggest a function name based on the similarity to known functions
			sort.Slice(knownFunctions, func(a, b int) bool {
				aSimilarity := similarity.Default(call.FunctionName, knownFunctions[a])
				bSimilarity := similarity.Default(call.FunctionName, knownFunctions[b])
				return aSimilarity > bSimilarity
			})

			if similarity.Default(call.FunctionName, knownFunctions[0]) > 0.9 {
				return compiler.Error(fmt.Sprintf("Unknown function '%s', did you mean '%s'?", call.FunctionName, knownFunctions[0]))
			}

			return compiler.Error(fmt.Sprintf("Unknown function '%s'", call.FunctionName))
		}

		// print builtin
		if call.FunctionName == "print" {
			parameters := call.Parameters
			expectedParameters := 1

			if len(parameters) < expectedParameters {
				return compiler.Error(fmt.Sprintf("Too few arguments in '%s' call", call.FunctionName))
			}

			if len(parameters) > expectedParameters {
				return compiler.Error(fmt.Sprintf("Too many arguments in '%s' call", call.FunctionName))
			}

			parameter := parameters[0][0]

			if parameter.Kind != token.Text {
				return compiler.Error(fmt.Sprintf("'%s' requires a text parameter instead of '%s'", call.FunctionName, string(parameter.Text)))
			}

			text := parameter.Text
			text = text[1 : len(text)-1]
			compiler.assembler.Println(string(text))
			compiler.funcCalls = compiler.funcCalls[:len(compiler.funcCalls)-1]
		}
	}

	compiler.Tokens = append(compiler.Tokens, t)
	return nil
}
