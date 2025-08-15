package tokenizer

import "testing"

func TestExtractLowerLetters(t *testing.T) {
  testCases := []struct {
    input string
    expected []rune
  }{
    {
      input: "Go is Fun!",
      expected: []rune{'o', 'i', 's', 'u', 'n'},
    },
    {
      input: "ABCdef",
      expected: []rune{'d', 'e', 'f'}, 
    }, 
    {
      input: "",
      expected: []rune{}, 
    },
    {
      input: "123?!",
      expected: []rune{},
    },
    {
      input: "n1g3l",
      expected: []rune{'n', 'g', 'l'},
    },
  }

  for _, tc := range testCases {
    result := ExtractLowerLetters(tc.input)

    if len(result) != len(tc.expected) {
      t.Fatalf("expected length %d, got %d", len(tc.expected), len(result))
    }

    for i, r := range result {
      if r != tc.expected[i] {
          t.Errorf("expected %q at index %d, got %q", tc.expected[i], i, r)
      }
    }
  }
}
