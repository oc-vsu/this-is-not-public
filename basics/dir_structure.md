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


```
project
│   README.md
│   go.mod
│   go.sum      
└───cmd
│   └───subfolder1
│       │   main.go
└───internal
│   └───app
│   │   └───model
│   │       │   model.go
│   └───pkg
│       └───keycloakApi
│           │   ...
```