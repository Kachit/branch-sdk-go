# Branch.io SDK GO (Unofficial)
[![Build Status](https://app.travis-ci.com/Kachit/branch-sdk-go.svg?branch=master)](https://app.travis-ci.com/github/Kachit/branch-sdk-go)
[![Codecov](https://codecov.io/gh/Kachit/branch-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/branch-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kachit/branch-sdk-go)](https://goreportcard.com/report/github.com/kachit/branch-sdk-go)
[![Release](https://img.shields.io/github/v/release/Kachit/branch-sdk-go.svg)](https://github.com/Kachit/branch-sdk-go/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/branch-sdk-go/blob/master/LICENSE)

## Description
Unofficial Branch.io reports API Client for Go

## API documentation
https://help.branch.io/developers-hub/docs/daily-exports

## Installation
```shell
go get -u github.com/kachit/branch-sdk-go
```
## Usage
```go
package main

import (
    "fmt"
    "github.com/kachit/branch-sdk-go"
    "time"
)

func main(){
    // Create a client instance
    cfg := branchio.NewConfig("Your Branch key", "Your Branch secret key")
    client, err := branchio.NewClientFromConfig(cfg, nil)
    if err != nil {
        fmt.Printf("config parameter error " + err.Error())
        panic(err)
    }
}
```
### Get events ontology
```go
ctx := context.Background()
dt := time.Date(2022, 1, 30, 0, 0, 0, 0, time.Local)
result, response, err := client.Export().GetEventOntology(ctx, dt)
```