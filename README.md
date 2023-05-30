![Requirement](https://img.shields.io/badge/golang-1.15+-blue.svg)

<details>
<summary><b>Table of content</b></summary>

## Table of content
   * [Why](#why)
   * [Mission](#mission)
   * [How](#how)
   * [Getting Started](#getting-started)
   * [Building from Source](#building-from-source)
   * [Abstraction](#abstraction)

</details>

## Why

A “dual-track” approach is essential to digital transformation, pairing traditional, IT-led grand scale transformation with a new form of business-led “rapid cycle innovation”.

## Mission

Enable `citizen development`.

## How

Metadata driven architecture and `REPL` alike runtime engine.

In the past, we describe our business with code; with FlexDB, we describe our business with metadata and FlexDB translates metadata into runtime on the fly.

```
Requirement ------------->  code  -> Business
          |
          V
Requirement -> metadata -> [code] -> Business
```

## Getting Started

Try it out first (requires docker...)

``` bash
docker run -d --name flexsql -e MYSQL_DATABASE=easyapp -e MYSQL_ALLOW_EMPTY_PASSWORD=1 mysql:5.7
docker run -it --user 10001 -p 10001:8000 dddplus/flexdb:latest
```

## Building from Source

``` bash
go install github.com/jteeuwen/go-bindata/go-bindata

go get github.com/agile-app/flexdb
cd $GOPATH/src/github.com/agile-app/flexdb
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com
make build
```

## Abstraction

```
Org(unit of tenant)
 ├── Quota 
 ├── Authz
 ├── API
 │    ├── Metadata API
 │    │     └── odata
 │    ├── Data API
 │    │     └── odata
 └── App
      ├── Model
      │   ├── Column
      │   │    └── Plugin (ColumnKind)
      │   │         ├── Validation
      │   │         ├── Indexer
      │   │         ├── Formula
      │   │         ├── CellValueGenerator
      │   │         └── View related
      │   ├── Feature
      │   ├── Trigger (context)
      │   │    ├── BeforeInsert(action)
      │   │    └── AfterInsert(action)
      │   └── Data
      │        ├── Row
      │        ├── Clob
      │        └── Index
      │             ├── StringIndex
      │             ├── NumberIndex
      │             ├── SpatialIndex
      │             └── TimeIndex
      ├── UI
      │   ├── Chart
      │   └── Page
      │        └── View
      ├── Share
      ├── Workflow
      ├── Connector
      ├── Automation
      └── Addons
          ├── Metrics
          └── Logs
```

## Inspired By

Salesforce, Odoo
