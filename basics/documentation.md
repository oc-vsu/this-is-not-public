---
layout: default
title: Dokumentation
parent: Basics
nav_order: 5
---

# Dokumentation

- Kommentare direkt über:
    - type
    - variable
    - constant 
    - function
- Package Dokumentation kommen in eigene Datei *doc.go*


```go
// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
```


## Tool
- Integriertes Dokumentationstool
- liest "normale" Kommentare aus
- kann lokalen Webserver hochfahren um Dokumentation als Website darzustellen:   
    `godoc -http=:6060`
    - über die Flag `-index` wird ein Suchindex aufgebaut

## Links
- [godoc](https://go.dev/blog/godoc)
- [godoc cmd tool](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)
- [godoc guide/example auf github ](https://github.com/natefinch/godocgo)

