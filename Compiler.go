package main

import (
	"fmt"
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
	var buffer [4096]byte

	for {
		n, err := file.Read(buffer[:])

		if n > 0 {
			compiler.processBuffer(buffer[:n])
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

func (compiler *Compiler) processBuffer(buffer []byte) {
	var (
		i          int
		c          byte
		tokenStart int
	)

	for {
		c = buffer[i]

		// Identifiers
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			tokenStart = i

			for (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
				i++

				if i >= len(buffer) {
					return
				}

				c = buffer[i]
			}

			compiler.processToken(TokenIdentifier, buffer[tokenStart:i])
			continue
		}

		// Texts
		if c == '"' {
			tokenStart = i

			for {
				i++

				if i >= len(buffer) {
					return
				}

				c = buffer[i]

				if c == '"' {
					i++
					break
				}
			}

			compiler.processToken(TokenText, buffer[tokenStart:i])
			continue
		}

		// Bracket start
		if c == '(' {
			compiler.processToken(TokenBracketStart, buffer[i:i+1])
			i++
			continue
		}

		// Bracket end
		if c == ')' {
			compiler.processToken(TokenBracketEnd, buffer[i:i+1])
			i++
			continue
		}

		i++

		if i >= len(buffer) {
			return
		}
	}
}

func (compiler *Compiler) processToken(token Token, data []byte) {
	fmt.Println(token, string(data))

	switch token {
	case TokenIdentifier:
	}
}
