# go-check
Check equality

This package does not check absolute equality of objects. Rather it checks that
all the elements of the expected object exist in the given object.

## Example:

```go
type A struct {
  B *string
  C *string
}
d := "str"

expected := A{B:d&}
given := A{B:&d,C:&d}

b, err := CheckEquality(expected, given)
fmt.Println(b) // true
fmt.Println(err) // nil
```
