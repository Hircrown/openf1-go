# openf1-go

[![Go Mod](https://img.shields.io/github/go-mod/go-version/Hircrown/openf1-go)](go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/Hircrown/openf1-go)](https://goreportcard.com/report/github.com/Hircrown/openf1-go)

Unofficial [OpenF1](https://openf1.org) Go client library

## Purpose
`openf1-go` is a Go client library that provides a simple and consistent interface for interacting with all [OpenF1](https://openf1.org) API endpoints.  
It allows you to query race, driver, telemetry data, and much more using structured filters or built-in helper functions — without manually building URLs or handling HTTP requests.

## Installation

```bash
go get github.com/Hircrown/openf1-go
go mod tidy
```

## Example Code 
### Usage of a filter query and ValuesBetween

```go
package main

import (
	"fmt"
	"log"

	"github.com/Hircrown/openf1-go/openf1"
	"github.com/Hircrown/openf1-go/openf1/types"
)

func main() {
	client := openf1.NewClient()

	lapDuration, err := openf1.ValuesBetween(types.LapFilter{}, "LapDuration", "98", "98.5", true)
	if err != nil {
		log.Fatal(err)
	}

	laps, err := client.Laps(types.LapFilter{
		SessionKey:   "9888",
		DriverNumber: 166,
		LapDuration:  lapDuration,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, lap := range laps {
		fmt.Printf("Lap (%d) in %.3fs\n", lap.LapNumber, lap.LapDuration)
	}
}
```

### Usage of a built-in function
```go
package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Hircrown/openf1-go/openf1"
	"github.com/Hircrown/openf1-go/openf1/types"
)

func main() {
	client := openf1.NewClient()

	grid, err := client.StartingGridGP("Monza", "2025", false)
	if err != nil {
		log.Fatal(err)
	}

	for _, position := range grid {
		// to avoid api request limit
		time.Sleep(300 * time.Millisecond)
		driver, err := client.Driver(types.DriverFilter{
			DriverNumber: position.DriverNumber,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(
			"#%02d %s%s- Time: %.3fs\n",
			position.Position,
			driver[0].FullName,
			strings.Repeat(" ", 18-len(driver[0].FullName)),
			position.LapDuration,
		)
	}
}
```

## Notice
openf1-go is unofficial and is not associated in any way with the Formula 1 companies. F1, FORMULA ONE, FORMULA 1, FIA FORMULA ONE WORLD CHAMPIONSHIP, GRAND PRIX and related marks are trademarks of Formula One Licensing B.V.