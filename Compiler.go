package main

import (
	"io"
	"os"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/elf"
)

type Compiler struct {
	assembler *asm.Assembler
}

func NewCompiler() *Compiler {
	return &Compiler{
		assembler: asm.New(),
	}
}

// Compile reads the input file and generates an executable binary.
func (compiler *Compiler) Compile(inputFile string, outputFile string) error {
	file, err := os.Open(inputFile)

	if err != nil {
		return err
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
			break
		}

		return err
	}

	// Programs should always exit
	compiler.assembler.Exit(0)

	// Produce ELF binary
	binary := elf.New(compiler.assembler)
	return binary.WriteToFile(outputFile)
}

// processBuffer processes the partial read and returns how many bytes were processed.
// The remaining bytes will be prepended to the next call.
func (compiler *Compiler) processBuffer(buffer []byte) int {
	var (
		i              int
		c              byte
		processedBytes int
	)

	for i < len(buffer) {
		c = buffer[i]

		// Identifiers
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			processedBytes = i

			for (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
				i++

				if i >= len(buffer) {
					return processedBytes
				}

				c = buffer[i]
			}

			token := Token{TokenIdentifier, buffer[processedBytes:i]}
			compiler.processToken(token)
			processedBytes = i + 1
			continue
		}

		// Texts
		if c == '"' {
			processedBytes = i

			for {
				i++

				if i >= len(buffer) {
					return processedBytes
				}

				c = buffer[i]

				if c == '"' {
					i++
					break
				}
			}

			token := Token{TokenText, buffer[processedBytes:i]}
			compiler.processToken(token)
			processedBytes = i + 1
			continue
		}

		// Bracket start
		if c == '(' {
			token := Token{TokenBracketStart, buffer[i : i+1]}
			compiler.processToken(token)
			i++
			processedBytes = i
			continue
		}

		// Bracket end
		if c == ')' {
			token := Token{TokenBracketEnd, buffer[i : i+1]}
			compiler.processToken(token)
			i++
			processedBytes = i
			continue
		}

		// End of line
		if c == '\n' {
			token := Token{TokenEndOfLine, nil}
			compiler.processToken(token)
			i++
			processedBytes = i
			continue
		}

		i++
	}

	return processedBytes
}

// processToken processes a single token.
func (compiler *Compiler) processToken(token Token) {
	switch token.Kind {
	case TokenText:
		compiler.assembler.Println(string(token.Text[1 : len(token.Text)-1]))
	}
}
