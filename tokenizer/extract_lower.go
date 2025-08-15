package tokenizer

func ExtractLowerLetters(input string) []rune {
  var result []rune
  for _, char := range input {
    if char >= 'a' && char <= 'z' {
      result = append(result, char)
    }
  }
  return result
}
