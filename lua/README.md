# InkspanDocument Lua SDK



The Lua SDK for the InkspanDocument API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:Checkout()` — each with the same small set of operations (`load`, `create`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Tags](https://github.com/voxgig-sdk/inkspan-document-sdk/tags)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("inkspan-document_sdk")

local client = sdk.new()
```

### 4. Create, update, and remove

```lua
-- Create
local created, err = client:Checkout():create({ plan = "example_plan" })
if err then error(err) end

```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local health, err = client:Health():load()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Health():load()
-- result is the entity; data_get() reads its mock record; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
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
cd lua && busted test/
```


## Reference

### InkspanDocumentSDK

```lua
local sdk = require("inkspan-document_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### InkspanDocumentSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
| `Checkout` | `(data) -> CheckoutEntity` | Create a Checkout entity instance. |
| `Compress` | `(data) -> CompressEntity` | Create a Compress entity instance. |
| `FillForm` | `(data) -> FillFormEntity` | Create a FillForm entity instance. |
| `Health` | `(data) -> HealthEntity` | Create a Health entity instance. |
| `Html` | `(data) -> HtmlEntity` | Create a Html entity instance. |
| `Inbound` | `(data) -> InboundEntity` | Create an Inbound entity instance. |
| `Markdown` | `(data) -> MarkdownEntity` | Create a Markdown entity instance. |
| `Men` | `(data) -> MenEntity` | Create a Men entity instance. |
| `Merge` | `(data) -> MergeEntity` | Create a Merge entity instance. |
| `Portal` | `(data) -> PortalEntity` | Create a Portal entity instance. |
| `RootGet` | `(data) -> RootGetEntity` | Create a RootGet entity instance. |
| `Signup` | `(data) -> SignupEntity` | Create a Signup entity instance. |
| `Split` | `(data) -> SplitEntity` | Create a Split entity instance. |
| `Table` | `(data) -> TableEntity` | Create a Table entity instance. |
| `Template` | `(data) -> TemplateEntity` | Create a Template entity instance. |
| `Text` | `(data) -> TextEntity` | Create a Text entity instance. |
| `ToImage` | `(data) -> ToImageEntity` | Create a ToImage entity instance. |
| `Translate` | `(data) -> TranslateEntity` | Create a Translate entity instance. |
| `Watermark` | `(data) -> WatermarkEntity` | Create a Watermark entity instance. |
| `Webhook` | `(data) -> WebhookEntity` | Create a Webhook entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria, and return it. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity, and return it. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the entity
itself — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` / `create` | the entity, whose `data_get()` reads its record (a `table`) |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local health, err = client:Health():load()
    if err then error(err) end
    -- health is the loaded entity

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

### Entities

#### Checkout

| Field | Description |
| --- | --- |

Operations: Create.

API path: `/v1/billing/checkout`

#### Compress

| Field | Description |
| --- | --- |
| `file` |  |

Operations: Create.

API path: `/v1/pdf/compress`

#### FillForm

| Field | Description |
| --- | --- |
| `fields` |  |
| `file` |  |

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
| `filename` |  |
| `html` |  |

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
| `filename` |  |
| `markdown` |  |

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
| `files` |  |

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
| `email` |  |

Operations: Create.

API path: `/v1/signup`

#### Split

| Field | Description |
| --- | --- |
| `file` |  |

Operations: Create.

API path: `/v1/pdf/split`

#### Table

| Field | Description |
| --- | --- |
| `file` |  |

Operations: Create.

API path: `/v1/extract/tables`

#### Template

| Field | Description |
| --- | --- |
| `data` |  |
| `filename` |  |
| `template` |  |

Operations: Create.

API path: `/v1/generate/template`

#### Text

| Field | Description |
| --- | --- |
| `file` |  |

Operations: Create.

API path: `/v1/extract/text`

#### ToImage

| Field | Description |
| --- | --- |
| `dpi` |  |
| `file` |  |

Operations: Create.

API path: `/v1/pdf/to-images`

#### Translate

| Field | Description |
| --- | --- |
| `file` |  |
| `output` |  |
| `target_language` |  |
| `text` |  |

Operations: Create.

API path: `/v1/translate`

#### Watermark

| Field | Description |
| --- | --- |
| `file` |  |
| `text` |  |

Operations: Create.

API path: `/v1/pdf/watermark`

#### Webhook

| Field | Description |
| --- | --- |

Operations: Create.

API path: `/v1/billing/webhook`



## Entities


### Checkout

Create an instance: `local checkout = client:Checkout(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```lua
local checkout, err = client:Checkout():create({
  plan = "example_plan", -- string
})
```


### Compress

Create an instance: `local compress = client:Compress(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```lua
local compress, err = client:Compress():create({
  file = "example_file", -- string
})
```


### FillForm

Create an instance: `local fill_form = client:FillForm(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `fields` | `string` |  |
| `file` | `string` |  |

#### Example: Create

```lua
local fill_form, err = client:FillForm():create({
  fields = "example_fields", -- string
  file = "example_file", -- string
})
```


### Health

Create an instance: `local health = client:Health(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```lua
local health, err = client:Health():load()
```


### Html

Create an instance: `local html = client:Html(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `filename` | `string` |  |
| `html` | `string` |  |

#### Example: Create

```lua
local html, err = client:Html():create({
  html = "example_html", -- string
})
```


### Inbound

Create an instance: `local inbound = client:Inbound(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```lua
local inbound, err = client:Inbound():create({
})
```


### Markdown

Create an instance: `local markdown = client:Markdown(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `filename` | `string` |  |
| `markdown` | `string` |  |

#### Example: Create

```lua
local markdown, err = client:Markdown():create({
  markdown = "example_markdown", -- string
})
```


### Men

Create an instance: `local men = client:Men(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```lua
local men, err = client:Men():load()
```


### Merge

Create an instance: `local merge = client:Merge(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `files` | `table` |  |

#### Example: Create

```lua
local merge, err = client:Merge():create({
  files = {}, -- table
})
```


### Portal

Create an instance: `local portal = client:Portal(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```lua
local portal, err = client:Portal():create({
})
```


### RootGet

Create an instance: `local root_get = client:RootGet(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```lua
local root_get, err = client:RootGet():load()
```


### Signup

Create an instance: `local signup = client:Signup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` |  |

#### Example: Create

```lua
local signup, err = client:Signup():create({
  email = "example_email", -- string
})
```


### Split

Create an instance: `local split = client:Split(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```lua
local split, err = client:Split():create({
  file = "example_file", -- string
})
```


### Table

Create an instance: `local table = client:Table(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```lua
local table, err = client:Table():create({
  file = "example_file", -- string
})
```


### Template

Create an instance: `local template = client:Template(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `table` |  |
| `filename` | `string` |  |
| `template` | `string` |  |

#### Example: Create

```lua
local template, err = client:Template():create({
  data = {}, -- table
  template = "example_template", -- string
})
```


### Text

Create an instance: `local text = client:Text(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```lua
local text, err = client:Text():create({
  file = "example_file", -- string
})
```


### ToImage

Create an instance: `local to_image = client:ToImage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dpi` | `number` |  |
| `file` | `string` |  |

#### Example: Create

```lua
local to_image, err = client:ToImage():create({
  file = "example_file", -- string
})
```


### Translate

Create an instance: `local translate = client:Translate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `any` |  |
| `output` | `string` |  |
| `target_language` | `string` |  |
| `text` | `any` |  |

#### Example: Create

```lua
local translate, err = client:Translate():create({
  target_language = "example_target_language", -- string
})
```


### Watermark

Create an instance: `local watermark = client:Watermark(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |
| `text` | `string` |  |

#### Example: Create

```lua
local watermark, err = client:Watermark():create({
  file = "example_file", -- string
})
```


### Webhook

Create an instance: `local webhook = client:Webhook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```lua
local webhook, err = client:Webhook():create({
})
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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **RatelimitFeature**: Rate limiting
- **RetryFeature**: Retry
- **TestFeature**: Test transport
- **TimeoutFeature**: Timeout

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── inkspan-document_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── schema.lua               -- Generated option + entity specs
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`inkspan-document_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```lua
local health = client:Health()
health:load()

-- health:data_get() now returns the health data from the last load
-- health:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
