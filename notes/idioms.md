# Go Idioms & Patterns

A living list of Go language patterns, syntax, and habits I want to keep using in production code.

---

## Arrays vs Slices

- **Arrays**: fixed length, part of the type (`[3]int` is not `[4]int`)
- **Slices**: dynamic length, backed by an underlying array
- Prefer slices for most code — more flexible, functions can take `[]T` without knowing the size

## Pointers

> **Mental model:** A pointer is just a note with directions to where the thing lives, so if you hand someone the note, they can change the thing itself instead of a photocopy.

- A pointer holds the **memory address** of a value, not the value itself.
- Declared with `*T` (pointer to a value of type `T`).
- Get the address of a value with `&`.
- Dereference (access or change the value) with `*`.

```go
x := 42
p := &x        // p is *int (pointer to int)
fmt.Println(*p) // prints 42

*p = 21        // changes x to 21
fmt.Println(x) // prints 21

## Implementing fmt.Stringer for Custom Printing
- Any type with a `String() string` method automatically satisfies `fmt.Stringer`.
- Useful for:
  - Debugging (custom format output)
  - Logging clean, readable values
  - Printing without exposing internal structure
```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d years)", p.Name, p.Age)
}

