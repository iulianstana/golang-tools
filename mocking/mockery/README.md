# Mocking using mockery package

[![Source code](../../src/reference.png)](https://github.com/vektra/mockery)


## Installation

`go get github.com/vektra/mockery/.../`, then `$GOPATH/bin/mockery`


## Example

Run: `mockery -name=HashTable` and the a mock will be output to `mocks/HashTable.go`

With the generated mock, new test can be build and methods output can be alternate.


## Return Value Provider Functions

If your tests need access to the arguments to calculate the return values,
set the return value to a function that takes the method's arguments as its own
arguments and returns the return value. For example, given this interface:

```go
package test

type Proxy interface {
  passthrough(s string) string
}
```

The argument can be passed through as the return value:

```go
import . "github.com/stretchr/testify/mock"

Mock.On("passthrough", AnythingOfType("string")).Return(func(s string) string {
    return s
})
```
