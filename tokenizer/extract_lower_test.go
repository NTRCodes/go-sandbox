package tokenizer

import "testing"

func TestExtractLowerLetters(t *testing.T) {
  input := "Go is Fun!"
  expected := []rune{'o', 'i', 's', 'u', 'n'}

  result := ExtractLowerLetters(input)

  if len(result) != len(expected) {
    t.Fatalf("expected length %d, got %d", len(expected), len(result))
  }

  for i := range result {
    if result[i] != expected[i] {
        t.Errorf("expected %q at index %d, got %q", expected[i], i, result[i])
    }
  }
}
