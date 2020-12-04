# netpalmgo
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/0f2e53ed370844ff8696317a51be1e9e)](https://app.codacy.com/gh/jpbede/netpalmgo?utm_source=github.com&utm_medium=referral&utm_content=jpbede/netpalmgo&utm_campaign=Badge_Grade)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/jpbede/netpalmgo)](https://pkg.go.dev/github.com/jpbede/netpalmgo)
![test](https://github.com/jpbede/netpalmgo/workflows/test/badge.svg)
[![codecov](https://codecov.io/gh/jpbede/netpalmgo/branch/main/graph/badge.svg)](https://codecov.io/gh/jpbede/netpalmgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/jpbede/netpalmgo)](https://goreportcard.com/report/github.com/jpbede/netpalmgo)

Go package for [netpalm API](https://github.com/tbotnz/netpalm)

## netpalmgo support

I maintain a community on the networktocode slack channel

#netpalmgo on networktocode.slack.com

## Examples

### Queueing a get task

```go
package main

import (
  "context"
  "github.com/jpbede/netpalmgo"
  "github.com/jpbede/netpalmgo/models"
)

func main() {
  nc := netpalmgo.New("https://netapi.mynet.net", "<apikey>")

  m := models.GetConfigRequest{
    Library: models.LibraryNetmiko,
    ConnectionArgs: models.ConnectionArgs{
      DeviceType: "cisco_ios",
      Host: "10.10.10.1",
      Username: "admin",
      Password: "abc123",
    },
    QueueStrategy: models.QueueStrategyFIFO,
    Command: []string{"show ip bgp sum"},
  }

  // send task
  d, err := nc.GetConfig().WithRequest(context.Background(), m)
  if err != nil {
    panic(err)
  }

  // wait blocking for task to finish
  d, err = nc.WaitForResult(context.Background(), d)
  if err != nil {
    panic(err)
  }
}
```

### Queueing a set command

```go
package main

import (
  "context"
  "github.com/jpbede/netpalmgo"
  "github.com/jpbede/netpalmgo/models"
)

func main() {
  nc := netpalmgo.New("https://netapi.mynet.net", "<apikey>")

  m := models.SetConfigRequest{
    Library: models.LibraryNetmiko,
    ConnectionArgs: models.ConnectionArgs{
      DeviceType: "cisco_ios",
      Host: "10.10.10.1",
      Username: "admin",
      Password: "abc123",
    },
    QueueStrategy: models.QueueStrategyFIFO,
    Config: []string{
      "set int tunnel tun0 remote-ip 192.168.2.1",
    },
  }

  // send set task
  d, err := nc.SetConfig().Set(context.Background(), false, m)
  if err != nil {
    panic(err)
  }

  // wait blocking for task to finish
  d, err = nc.WaitForResult(context.Background(), d)
  if err != nil {
    panic(err)
  }
}
```