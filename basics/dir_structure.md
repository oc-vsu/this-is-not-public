---
layout: default
title: Struktur
parent: Basics
nav_order: 3
---

# Datei & Ordner Struktur

- Eine offiziell definierte Projekt Struktur gibt es nicht
- "industrie Standard" folgt https://github.com/golang-standards/project-layout, darunter u.a.:
  - https://github.com/kubernetes/kubernetes
  - https://github.com/moby/moby
  - https://github.com/hashicorp/terraform

Beispiel:
```
project
│   README.md
│   go.mod
│   go.sum      
└───cmd
│   └───cmd1
│       │   main.go
│   └───cmd2
│       │   main.go
└───internal
│   └───app
│   │   └───cmd1
│   │   │   └───foo
│   │   │       │ ....
│   │   └───cmd2
│   │   │   └───bar
│               │ ....
│   └───pkg
│       └───keycloakApi
│           │   ...
```
- der pkg Ordner wird für Packages verwendet, welche ohne Abhängigkeiten von allen Subprogrammen in einem Projekt verwendet werden kann:
  - cmd1 und cmd1 können jeweils auf das keycloakApi zugreifen
  - cmd1 darf nicht auf cmd2/bar und cmd2 nicht auf cmd1/foo zugreifen