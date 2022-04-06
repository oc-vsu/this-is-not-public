---
layout: default
title: Setup
parent: Basics
nav_order: 1
---

# Setup

Anweisungen unter: [https://go.dev/doc/install](https://go.dev/doc/install)

## GOROOT & GOPATH

- GOROOT ist der Installationspfad unter dem GO installiert ist:
    ```
    $ go env GOROOT
    /usr/local/go
    ```
- GOPATH definiert den Workspace unter dem GO Projekte erstellt werden. I.d.R im home-Verzeichnis:
    ```
    $ go env GOPATH
    /Users/vsu/go
    ```

### Links
- [gopath env var](https://pkg.go.dev/cmd/go#hdr-GOPATH_environment_variable)


## Go Modules
- Dependency management
- wurde in v1.11 eingeführt
- löst die Entwicklung in einem GOPATH ab
    - GOPATH wird weiterhin von der Runtime als Download Verzeichnis für Packages verwendet
- `go mod tidy` aktualisiert die go.mod Datei
    - fügt fehlende Packages hinzu
    - entfernt nicht benutzte 
- die go.sum Datei enthält die Checksum der installierten Packages

```go
// Beispiel 1
module foo

go 1.18 // optional


// Beispiel 2
module opitz-consulting.com/project/project

go 1.18 // optional
```

### Links
- [gopath and modules](https://pkg.go.dev/cmd/go#hdr-GOPATH_and_Modules)
- [modules documentation](https://go.dev/ref/mod)
- [the go.mod file](https://pkg.go.dev/cmd/go#hdr-The_go_mod_file)
