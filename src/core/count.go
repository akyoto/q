package core

type counter = uint8

// count stores how often a certain statement appeared so we can generate a unique label from it.
type count struct {
	data counter
}