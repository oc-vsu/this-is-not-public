---
layout: default
title: Unit - Testing
parent: Testing
nav_order: 1
---

# Unit Testing & Mocking

## Unit Testing

- Test files müssen mit `_test.go` enden:
    ```
    project
    └───example
       │    example.go
       │    example_test.go
    ```
- Tests verwenden das "testing" package
- Methoden Signatur:
    - "Test" Prefix
    - `t *testing.T` muss als Argument verwendet werden
- für Assertion [github.com/stretchr/testify](https://github.com/stretchr/testify#assert-package)
- für das Ausführen der Tests:
    - `go test [package]`
    - `go test ./...`

```go
package example

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
    var a = 123
    var b = 123

    assert.Equal(t, a, b)

    a = a + 1

    assert.NotEqual(t, a, b)

    t.Run("sub_test", func(t *testing.T) {
        if true != true {
            t.Errorf("test failed")
        }

        t.Log("test logging")
    })
}
```

- für zusätzliches Setup vor Ausführung der Tests, kann TestMain verwendet werden:
    ```go
    func TestMain(m *testing.M)
    ```

## Links
- [github.com/stretchr/testify#assert-package](https://github.com/stretchr/testify#assert-package)
- [go.dev/doc/tutorial/add-a-test](https://go.dev/doc/tutorial/add-a-test)
- [pkg.go.dev/testing](https://pkg.go.dev/testing)