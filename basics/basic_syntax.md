---
layout: default
title: Basic Syntax
parent: Basics
nav_order: 2
---

# Basic Syntax

- main funktion im main package ist Einstiegspunkt
	```go
	package main

	func main() {
		
	}
	```
- keine nullable types, Pointer können als Ersatz fungieren
- keine Keywords für Sichtbarkeit
    - camelCase: private
    - PascalCase: public (exported)
	```go
	// private
	func fooBar() {}
	var foo string
	const foo string

	// public (exported)
	func FooBar() {}
	var Foo string
	const Foo string
	```
- multiple Rückgabewerte
- keine Exceptions dafür error type
	```go
	package main

	import "error"

	func fooBar() (string, error) {
		return "fooBar", errors.New("error")
	}
	```
- mit dem "blank identifier", können ungenutzte Werte verworfen werden, vergleichbar mit /dev/null
	```go
	package main

	import "strconv"

	func fooBar() (string, error) {
		return "fooBar", errors.New("error")
	}

	func main() {
		_, err := fooBar()
	
	}
	```

## "Objektorientierung"
- Go ist prozedural
    - allerdings gibt es 
- keine Klassen dafür Structs
- Funktionen können auf Structs gebindet werden
- "New" Funktion als Konstruktor
- interfaces
    - 

```go
// example/example.go
package example

type ExampleIFace interface {
	FooBar() string
}

type Example struct {
	Value string
}

func New(v string) *Example {
	return &Example{
		Value: v,
	}
}

func (e *Example) FooBar() string {
	return e.Value
}

// main.go
package main

import (
	"example"
	"fmt"
)

func main() {
	var e example.ExampleIFace = example.New("foo")

	fmt.Println(e.FooBar())

    // will output "foo"
}
```

# go fmt your code!
- integriertes Formatierungs - Tool [go.dev/blog/gofmt]:(https://go.dev/blog/gofmt)
	- `go fmt [package]`
	- `go fmt ./...` 