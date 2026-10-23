# InkspanDocument Golang SDK



The Golang SDK for the InkspanDocument API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Checkout(nil)` — each with the same small set of operations (`Load`, `Create`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/inkspan-document-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Tags](https://github.com/voxgig-sdk/inkspan-document-sdk/tags) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/inkspan-document-sdk/go=../inkspan-document-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the entity, and for
`List` a `[]any` of entities, one per record (there is no `{ok, data}`
wrapper), so check `err` and read a record through the entity's
`Data()`.

```go
package main

import (
    "fmt"
    sdk "github.com/voxgig-sdk/inkspan-document-sdk/go"
)

func main() {
    client := sdk.New()

    // Create a checkout.
    created, err := client.Checkout(nil).Create(map[string]any{"plan": "example_plan"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created.(sdk.Entity).Data())
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
health, err := client.Health(nil).Load(nil, nil)
if err != nil {
    // handle err
    return
}
_ = health
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

health, err := client.Health(nil).Load(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(health.(sdk.Entity).Data()) // the entity's mock record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewInkspanDocumentSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
INKSPAN_DOCUMENT_TEST_LIVE=TRUE
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewInkspanDocumentSDK

```go
func NewInkspanDocumentSDK(options map[string]any) *InkspanDocumentSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *InkspanDocumentSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### InkspanDocumentSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Checkout` | `(data map[string]any) InkspanDocumentEntity` | Create a Checkout entity instance. |
| `Compress` | `(data map[string]any) InkspanDocumentEntity` | Create a Compress entity instance. |
| `FillForm` | `(data map[string]any) InkspanDocumentEntity` | Create a FillForm entity instance. |
| `Health` | `(data map[string]any) InkspanDocumentEntity` | Create a Health entity instance. |
| `Html` | `(data map[string]any) InkspanDocumentEntity` | Create a Html entity instance. |
| `Inbound` | `(data map[string]any) InkspanDocumentEntity` | Create an Inbound entity instance. |
| `Markdown` | `(data map[string]any) InkspanDocumentEntity` | Create a Markdown entity instance. |
| `Men` | `(data map[string]any) InkspanDocumentEntity` | Create a Men entity instance. |
| `Merge` | `(data map[string]any) InkspanDocumentEntity` | Create a Merge entity instance. |
| `Portal` | `(data map[string]any) InkspanDocumentEntity` | Create a Portal entity instance. |
| `RootGet` | `(data map[string]any) InkspanDocumentEntity` | Create a RootGet entity instance. |
| `Signup` | `(data map[string]any) InkspanDocumentEntity` | Create a Signup entity instance. |
| `Split` | `(data map[string]any) InkspanDocumentEntity` | Create a Split entity instance. |
| `Table` | `(data map[string]any) InkspanDocumentEntity` | Create a Table entity instance. |
| `Template` | `(data map[string]any) InkspanDocumentEntity` | Create a Template entity instance. |
| `Text` | `(data map[string]any) InkspanDocumentEntity` | Create a Text entity instance. |
| `ToImage` | `(data map[string]any) InkspanDocumentEntity` | Create a ToImage entity instance. |
| `Translate` | `(data map[string]any) InkspanDocumentEntity` | Create a Translate entity instance. |
| `Watermark` | `(data map[string]any) InkspanDocumentEntity` | Create a Watermark entity instance. |
| `Webhook` | `(data map[string]any) InkspanDocumentEntity` | Create a Webhook entity instance. |

### Entity interface (InkspanDocumentEntity)

All entities implement the `InkspanDocumentEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria, and return it. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity, and return it. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the entity
itself — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` | the entity, whose `Data()` reads its record (`map[string]any`) |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    checkout, err := client.Checkout(nil).Create(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // checkout is the entity; checkout.(sdk.Entity).Data() reads its record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### Checkout

| Field | Description |
| --- | --- |

Operations: Create.

API path: `/v1/billing/checkout`

#### Compress

| Field | Description |
| --- | --- |
| `"file"` |  |

Operations: Create.

API path: `/v1/pdf/compress`

#### FillForm

| Field | Description |
| --- | --- |
| `"fields"` |  |
| `"file"` |  |

Operations: Create.

API path: `/v1/pdf/fill-form`

#### Health

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/health`

#### Html

| Field | Description |
| --- | --- |
| `"filename"` |  |
| `"html"` |  |

Operations: Create.

API path: `/v1/generate/html`

#### Inbound

| Field | Description |
| --- | --- |

Operations: Create.

API path: `/webhooks/inbound`

#### Markdown

| Field | Description |
| --- | --- |
| `"filename"` |  |
| `"markdown"` |  |

Operations: Create.

API path: `/v1/generate/markdown`

#### Men

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/v1/me`

#### Merge

| Field | Description |
| --- | --- |
| `"files"` |  |

Operations: Create.

API path: `/v1/pdf/merge`

#### Portal

| Field | Description |
| --- | --- |

Operations: Create.

API path: `/v1/billing/portal`

#### RootGet

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/`

#### Signup

| Field | Description |
| --- | --- |
| `"email"` |  |

Operations: Create.

API path: `/v1/signup`

#### Split

| Field | Description |
| --- | --- |
| `"file"` |  |

Operations: Create.

API path: `/v1/pdf/split`

#### Table

| Field | Description |
| --- | --- |
| `"file"` |  |

Operations: Create.

API path: `/v1/extract/tables`

#### Template

| Field | Description |
| --- | --- |
| `"data"` |  |
| `"filename"` |  |
| `"template"` |  |

Operations: Create.

API path: `/v1/generate/template`

#### Text

| Field | Description |
| --- | --- |
| `"file"` |  |

Operations: Create.

API path: `/v1/extract/text`

#### ToImage

| Field | Description |
| --- | --- |
| `"dpi"` |  |
| `"file"` |  |

Operations: Create.

API path: `/v1/pdf/to-images`

#### Translate

| Field | Description |
| --- | --- |
| `"file"` |  |
| `"output"` |  |
| `"target_language"` |  |
| `"text"` |  |

Operations: Create.

API path: `/v1/translate`

#### Watermark

| Field | Description |
| --- | --- |
| `"file"` |  |
| `"text"` |  |

Operations: Create.

API path: `/v1/pdf/watermark`

#### Webhook

| Field | Description |
| --- | --- |

Operations: Create.

API path: `/v1/billing/webhook`



## Entities


### Checkout

Create an instance: `checkout := client.Checkout(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Example: Create

```go
result, err := client.Checkout(nil).Create(map[string]any{
    "plan": "example_plan",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Compress

Create an instance: `compress := client.Compress(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```go
result, err := client.Compress(nil).Create(map[string]any{
    "file": "example_file",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### FillForm

Create an instance: `fillForm := client.FillForm(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `fields` | `string` |  |
| `file` | `string` |  |

#### Example: Create

```go
result, err := client.FillForm(nil).Create(map[string]any{
    "fields": "example_fields",
    "file": "example_file",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Health

Create an instance: `health := client.Health(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Example: Load

```go
health, err := client.Health(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(health.(sdk.Entity).Data()) // the loaded entity's record
```


### Html

Create an instance: `html := client.Html(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `filename` | `string` |  |
| `html` | `string` |  |

#### Example: Create

```go
result, err := client.Html(nil).Create(map[string]any{
    "html": "example_html",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Inbound

Create an instance: `inbound := client.Inbound(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Example: Create

```go
result, err := client.Inbound(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Markdown

Create an instance: `markdown := client.Markdown(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `filename` | `string` |  |
| `markdown` | `string` |  |

#### Example: Create

```go
result, err := client.Markdown(nil).Create(map[string]any{
    "markdown": "example_markdown",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Men

Create an instance: `men := client.Men(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Example: Load

```go
men, err := client.Men(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(men.(sdk.Entity).Data()) // the loaded entity's record
```


### Merge

Create an instance: `merge := client.Merge(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `files` | `[]any` |  |

#### Example: Create

```go
result, err := client.Merge(nil).Create(map[string]any{
    "files": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Portal

Create an instance: `portal := client.Portal(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Example: Create

```go
result, err := client.Portal(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### RootGet

Create an instance: `rootGet := client.RootGet(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Example: Load

```go
rootGet, err := client.RootGet(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(rootGet.(sdk.Entity).Data()) // the loaded entity's record
```


### Signup

Create an instance: `signup := client.Signup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` |  |

#### Example: Create

```go
result, err := client.Signup(nil).Create(map[string]any{
    "email": "example_email",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Split

Create an instance: `split := client.Split(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```go
result, err := client.Split(nil).Create(map[string]any{
    "file": "example_file",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Table

Create an instance: `table := client.Table(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```go
result, err := client.Table(nil).Create(map[string]any{
    "file": "example_file",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Template

Create an instance: `template := client.Template(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `map[string]any` |  |
| `filename` | `string` |  |
| `template` | `string` |  |

#### Example: Create

```go
result, err := client.Template(nil).Create(map[string]any{
    "data": map[string]any{},
    "template": "example_template",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Text

Create an instance: `text := client.Text(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```go
result, err := client.Text(nil).Create(map[string]any{
    "file": "example_file",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### ToImage

Create an instance: `toImage := client.ToImage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dpi` | `int` |  |
| `file` | `string` |  |

#### Example: Create

```go
result, err := client.ToImage(nil).Create(map[string]any{
    "file": "example_file",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Translate

Create an instance: `translate := client.Translate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `any` |  |
| `output` | `string` |  |
| `target_language` | `string` |  |
| `text` | `any` |  |

#### Example: Create

```go
result, err := client.Translate(nil).Create(map[string]any{
    "target_language": "example_target_language",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Watermark

Create an instance: `watermark := client.Watermark(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |
| `text` | `string` |  |

#### Example: Create

```go
result, err := client.Watermark(nil).Create(map[string]any{
    "file": "example_file",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Webhook

Create an instance: `webhook := client.Webhook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Example: Create

```go
result, err := client.Webhook(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```

## Features

This SDK ships 4 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`ratelimit`](#ratelimit) | Rate limiting |
| [`retry`](#retry) | Retry |
| [`test`](#test) | Test transport |
| [`timeout`](#timeout) | Timeout |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### ratelimit

Rate limiting.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Retry.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

Test transport.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Timeout.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **RatelimitFeature**: Rate limiting
- **RetryFeature**: Retry
- **TestFeature**: Test transport
- **TimeoutFeature**: Timeout

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

An operation returns the entity, and its `Data()` returns the record. Use
`core.ToMapAny()` to safely cast that record, or data nested in it, to
`map[string]any`: it returns `nil` for anything else, an entity included.

### Package structure

```
github.com/voxgig-sdk/inkspan-document-sdk/go/
├── inkspan-document.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/inkspan-document-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
health := client.Health(nil)
health.Load(nil, nil)

// health.Data() now returns the health data from the last load
// health.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
