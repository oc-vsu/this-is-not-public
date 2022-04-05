---
layout: default
title: Basic Syntax
parent: Basics
nav_order: 2
---

# Basic Syntax

- main packge + main funktion einstiegs punkt
- keine nullable types
    - pointer fungieren als Ersatz
- keine Keywords für Sichtbarkeit
    - camelCase: private
    - PascalCase: public
- multiple Rückgabewerte
- keine Exceptions dafür error type


```go
package main

func main() {
    var foo uint = 42


    fooShort := 42
}
```

## Objektorientierung
- keine Klassen dafür Structs
- Funktionen können auf Structs gebindet werden
- "New" Funktion als Konstruktor
- interfaces -> MOCKING?

```go
package main

func main() {
    var foo uint = 42


    fooShort := 42
}
```

# go fmt your code!
- integriertes formatierungs tool
https://go.dev/blog/gofmt