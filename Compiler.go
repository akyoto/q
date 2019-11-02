package main

import (
	"io"
	"os"

	"github.com/akyoto/asm"
)

type Compiler struct {
	assembler *asm.Assembler
}

func NewCompiler() *Compiler {
	return &Compiler{
		assembler: asm.New(),
	}
}

func (compiler *Compiler) Compile(inputFile string, outputFile string) {
	file, err := os.Open(inputFile)

	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	defer file.Close()

	var (
		buffer      [16384]byte
		unprocessed = make([]byte, 0, len(buffer))
		final       []byte
	)

	for {
		n, err := file.Read(buffer[:])

		if n > 0 {
			if len(unprocessed) > 0 {
				final = append(unprocessed, buffer[:n]...)
				unprocessed = unprocessed[:0]
			} else {
				final = buffer[:n]
			}

			processedBytes := compiler.processBuffer(final)

			if processedBytes < len(final) {
				unprocessed = append(unprocessed, final[processedBytes:]...)
			}
		}

		if err == nil {
			continue
		}

		if err == io.EOF {
			return
		}

		os.Stderr.WriteString(err.Error() + "\n")
		file.Close()
		os.Exit(1)
	}
}

func (compiler *Compiler) processBuffer(buffer []byte) int {
	var (
		i          int
		c          byte
		tokenStart int
	)

	for i < len(buffer) {
		c = buffer[i]

		// Identifiers
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			tokenStart = i

			for (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
				i++

				if i >= len(buffer) {
					return tokenStart
				}

				c = buffer[i]
			}

			token := Token{TokenIdentifier, buffer[tokenStart:i]}
			compiler.processToken(token)
			tokenStart = i + 1
			continue
		}

		// Texts
		if c == '"' {
			tokenStart = i

			for {
				i++

				if i >= len(buffer) {
					return tokenStart
				}

				c = buffer[i]

				if c == '"' {
					i++
					break
				}
			}

			token := Token{TokenText, buffer[tokenStart:i]}
			compiler.processToken(token)
			tokenStart = i + 1
			continue
		}

		// Bracket start
		if c == '(' {
			token := Token{TokenBracketStart, buffer[i : i+1]}
			compiler.processToken(token)
			i++
			tokenStart = i
			continue
		}

		// Bracket end
		if c == ')' {
			token := Token{TokenBracketEnd, buffer[i : i+1]}
			compiler.processToken(token)
			i++
			tokenStart = i
			continue
		}

		// End of line
		if c == '\n' {
			token := Token{TokenEndOfLine, nil}
			compiler.processToken(token)
			i++
			tokenStart = i
			continue
		}

		i++
	}

	return tokenStart
}

func (compiler *Compiler) processToken(token Token) {
	// fmt.Println(token.Kind, string(token.Text))

	switch token.Kind {
	case TokenIdentifier:
	}
}
