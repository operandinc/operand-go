# Go Operand

[![Go Reference](https://pkg.go.dev/badge/github.com/operandinc/operand-go)](https://pkg.go.dev/github.com/operandinc/operand-go)

The official [Operand](https://operand.ai) Go client library.

## Installation

Requires Go modules (e.g., `go.mod` file) & Go 1.18+.

```sh
go get -u github.com/operandinc/operand-go
```

Import the library:

```go
import "github.com/operandinc/operand-go"
```

## Usage

For complete API reference, please see our [API documentation](https://operand.ai/docs).

To index your first atom,

```go

client := operand.NewClient("<your-api-key>")

ctx := context.Background()

collection, err := client.GetCollection(ctx, "Discord Conversations")
if err != nil {
    // handle error
}

group, err := client.CreateGroup(ctx, &operand.CreateGroupRequest{
    CollectionID: collection.ID,
    Name: "DMs with Furqan",
})
if err != nil {
    // handle error
}

atom, err := client.CreateAtom(ctx, &operand.CreateAtomRequest{
    GroupID: group.ID,
    Content: "should be free up around 6",
    Properties: operand.Properties{
        "direction": "inbound",
    },
})
if err != nil {
    // handle error
}

// Done!

```
