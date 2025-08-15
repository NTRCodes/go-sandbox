package tokenizer

import "io"

type Tokenizer struct{
  input string
  pos int
}

func NewTokenizer(input string) *Tokenizer {
  return &Tokenizer{input: input}
}

func (t *Tokenizer) NextToken() (string, error) {
  return "", io.EOF
}
