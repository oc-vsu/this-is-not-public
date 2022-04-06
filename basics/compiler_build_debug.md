---
layout: default
title: Compiler, Build & Debug
parent: Basics
nav_order: 3
---

# Compiler, Build & Debug

<img src="../images/go_build.png" alt="building_gophers" width="200"/>

## Compiler
- Für Go gibt es 2 Compiler
    - gc
    - gccgo
- gc wurde ursprünglich in C geschrieben, seit Version 1.5 nur noch Go und Assembler 
    - gc ist der Standard-Compiler
- gccgo basiert auf GCC
    - compile Zeiten langsamer
    - größere Anzahl an unterstützen Prozessorarchitekturen


### Links
- [Release Notes v1.5](https://go.dev/doc/go1.5#c)
- [Unterschied gc / gccgo](https://stackoverflow.com/a/25811505)
- [gccgo setup](https://go.dev/doc/install/gccgo)

## Build
### Install command
- go install compiliert und installiert packages und dependencies
    - `go install [packages]`
    - `go install ./...`
- mehr details auf: [go.dev](https://go.dev/ref/mod#go-install)

### Build command
- go build compiliert die angegebenen packages / dateien.
    - mit dem Parameter "-o" kann ein Ausgabepfad angegeben werden    
    `go build -o my_binary main.go`
- cross-compiling, durch das Voranstellen der Env Variablen GOOS, GOARCH:
    - `GOOS=windows GOARCH=amd64 go build -o my_binary.exe main.go`
    - mit `go tool dist list` können verfügbare Platformen 
- mehr details auf [go.dev](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)


### Build constraints
- stehen am Anfang einer Go - Datei
    - `// +build <tag>`
- können Abhängigkeiten zu GOOS und GOARCH haben:
    - `// +build darwin,arm64` kompiliert nur auf MacOS, ARM64 _(siehe oben, cross-compiling)_
    - `// +build windows` kompiliert nur auf Windows _(siehe oben, cross-compiling)_
- eigene Tags möglich:
    - `// +build foo`, benötigt die Übergabe von `-tags foo` im build command.
        - `go build -o output -tags foo .`
- contraints können negiert werden:
    - `// +build !linux` // kompiliert nicht für Linux

### Run command
- go run compiliert und führt das angegbene package aus
    - `go run [packages]`   
    - `go run ./...`

### Links
- [compile & install an application](https://go.dev/doc/tutorial/compile-install)
- [build constraints](https://pkg.go.dev/go/build#hdr-Build_Constraints)
- [compile and run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)

## Debugging

### Visual Studio Code
<img src="../images/debugging/start_debugger.png" height="250">

A. main.go bzw. main package öffnen und auf "Ausführen und Debuggen" clicken _oder_   
B. launch.json erstellen und build Konfiguration anlegen:
- Go: launch file auswählen
- Konfiguration nach Belieben anpassen