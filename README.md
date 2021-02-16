# Branch.io SDK GO (Unofficial)
Golang SDK for Branch.io API (Unofficial)

## Description
Unofficial Branch.io API Client for Go

## API documentation
https://help.branch.io/developers-hub/docs/daily-exports

## Download
```shell
go get -u github.com/kachit/branch-sdk-go
```

```go
package main

import (
	"fmt"
	branchio_sdk "github.com/kachit/branchio-sdk-go"
	"time"
)

func main(){
        cfg := branchio_sdk.NewConfig("key_live", "secret_live")
        client := branchio_sdk.NewClientFromConfig(cfg, nil)
        now := time.Date(2021, 1, 30, 0, 0, 0, 0, time.Local)
        response, err := client.Exports().LinksList(now)
}
```
