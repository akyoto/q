package token

// Producer defines a token producer.
type Producer struct {
	Tokens []Token
}

// PreviousToken returns the last token with an offset.
func (producer *Producer) PreviousToken(offset int) Token {
	return producer.Tokens[len(producer.Tokens)+offset]
}
