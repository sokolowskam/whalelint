# <img width="22px" src="https://user-images.githubusercontent.com/5306361/110181582-6c807f80-7e0c-11eb-81c8-36d6a9c0db0b.png"> WhaleLint <img align="right" style="position: relative; top: 10px;" src="https://github.com/cremindes/whalelint/workflows/build/badge.svg" />

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

> *Disclaimer: this has started out as a pet-project while learning Golang.*

Dockerfile linter written in Go.

<p align="center">
  <img width="500px" src="docs/illustration/illustration.svg"/> 
</p>

## Sample output

<p align="center">
  <img width="750px" src="https://user-images.githubusercontent.com/5306361/110198673-775f0280-7e54-11eb-8e4e-ab6350fb4e7d.png"/>
</p>

## Rules

Each Dockerfile AST element has a corresponding set of rules. Click on the picture for details.

<p align="center"><a href="docs/rule/readme.md">
  <img width="500px" src="https://user-images.githubusercontent.com/5306361/110181292-bfa60280-7e0b-11eb-8437-d9ec9c45df62.png" />
</a/</p>

## Development

### Roadmap

| Feature |  | Status |
|---|---|---|
| Extendable ruleset|  | ![Done](https://img.shields.io/static/v1?label=&message=Done&color=Green) | 
| CLI |  |![Done](https://img.shields.io/static/v1?label=&message=Done&color=Green)  |  |
| Configurable Output | | ![Done](https://img.shields.io/static/v1?label=&message=Done&color=Green)
| - JSON | ![Done](https://img.shields.io/static/v1?label=&message=Done&color=Green) |
| - Colored Summary | ![Done](https://img.shields.io/static/v1?label=&message=Done&color=Green) |
| Rule pass | | ![NotYetStarted](https://img.shields.io/static/v1?label=&message=NoYetStarted&color=lightgrey) |
| - Per line | ![NotYetStarted](https://img.shields.io/static/v1?label=&message=NoYetStarted&color=lightgrey) |
| - Config file | ![NotYetStarted](https://img.shields.io/static/v1?label=&message=NoYetStarted&color=lightgrey) |
| Config file | | ![NotYetStarted](https://img.shields.io/static/v1?label=&message=NoYetStarted&color=lightgrey) |
| - Rule profiles | ![NotYetStarted](https://img.shields.io/static/v1?label=&message=NoYetStarted&color=lightgrey) |
| IDE plugins/extensions | | ![InProgress](https://img.shields.io/static/v1?label=&message=InProgress&color=blue)
| - VSCode | ![Preview](https://img.shields.io/static/v1?label=&message=Preview&color=blue)
| - JetBrains | ![InProgress](https://img.shields.io/static/v1?label=&message=InProgress&color=blue)

### Design Decisions

A collection of documents describing the thought process behind selected design decisions. [Link >](docs/design/readme.md)

### Contribution Guide

[Link > TODO](docs/contribution/readme.md)

## Plugins

### JetBrains

![InProgress](https://img.shields.io/static/v1?label=&message=InProgress&color=blue)

### VSCode

[![Version](https://vsmarketplacebadge.apphb.com/version/tamasgbarna.whalelint.svg)](https://marketplace.visualstudio.com/items?itemName=tamasgbarna.whalelint) 
[![Installs](https://vsmarketplacebadge.apphb.com/installs-short/tamasgbarna.whalelint.svg)](https://marketplace.visualstudio.com/items?itemName=tamasgbarna.whalelint)

<p align="center">
  <img src="https://user-images.githubusercontent.com/5306361/110014611-4c28c600-7d23-11eb-915d-114aca6470b2.gif">
</p>

## Alternatives

[Alternatives](docs/alternatives/readme.md)