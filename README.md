# Branch.io SDK GO (Unofficial)
[![Build Status](https://travis-ci.org/Kachit/branch-sdk-go.svg?branch=master)](https://travis-ci.org/Kachit/branch-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/branch-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/branch-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kachit/branch-sdk-go)](https://goreportcard.com/report/github.com/kachit/branch-sdk-go)
[![Release](https://img.shields.io/github/v/release/Kachit/branch-sdk-go.svg)](https://github.com/Kachit/branch-sdk-go/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/branch-sdk-go/blob/master/LICENSE)

## Description
Unofficial Branch.io API Client for Go

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
	branchio_sdk "github.com/kachit/branch-sdk-go"
	"time"
)

func main(){
        cfg := branchio_sdk.NewConfig("key_live", "secret_live")
        client := branchio_sdk.NewClientFromConfig(cfg, nil)
        dt := time.Date(2021, 1, 30, 0, 0, 0, 0, time.Local)
        response, err := client.Export().GetEventOntology(dt)
}
```
