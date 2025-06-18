# UUID

**UUID** This package is a utility where there is a UUID data type. It is useful when we have a database with UUID types, The purpose of this package is to avoid repetitive logic when generating, validating and parsing UUIDs. So is in charge of abstracting that part of the logic and allows us to manage more efficiently what has to do with uuids.
![Go](https://img.shields.io/badge/Golang-1.21-blue.svg?logo=go&longCache=true&style=flat)

This package makes use of the external package [google uuid](github.com/google/uuid).

## Getting Started

[Go](https://golang.org/) is required in version 1.21 or higher.

### Install

`go get -u github.com/pretty-technical-io/prettytechnical-tools-uuid`

### Features

* [x] **Lightweight**, less than 200 lines of code.
* [x] **Easy** to use.

### Basic use

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	uuid "github.com/pretty-technical-io/prettytechnical-tools-uuid"
)

type Player struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func main() {
	http.HandleFunc("/player", GetPLayerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetPLayerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	player := getPlayer(ctx)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(player)

}

func getPlayer(_ context.Context) *Player {
	// logic to get from database a player for example.
	id, err := uuid.New()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Player{
		ID:   id,
		Name: "player",
	}
}

```