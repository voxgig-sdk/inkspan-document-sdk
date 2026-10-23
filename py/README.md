# InkspanDocument Python SDK



The Python SDK for the InkspanDocument API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Checkout()` — each
carrying a small, uniform set of operations (`load`, `create`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Tags](https://github.com/voxgig-sdk/inkspan-document-sdk/tags)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
from inkspandocument_sdk import InkspanDocumentSDK

client = InkspanDocumentSDK()
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.Checkout().create({"plan": "example_plan"})

```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    health = client.Health().load()
    print(health.data_get())
except Exception as err:
    print(f"load failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = InkspanDocumentSDK.test()

# Entity ops return the entity, and list one per record; they raise on error.
health = client.Health().load()
# data_get() on an entity reads its mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = InkspanDocumentSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### InkspanDocumentSDK

```python
from inkspandocument_sdk import InkspanDocumentSDK

client = InkspanDocumentSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = InkspanDocumentSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### InkspanDocumentSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria, and return it. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity, and return it. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the entity, and `list` a `list` of entities, one
per record; an entity's `data_get()` reads its record (a `dict`). They raise
on error, so wrap calls in `try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `checkout = client.Checkout()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```python
checkout = client.Checkout().create({
    "plan": "example_plan",  # str
})
```


### Compress

Create an instance: `compress = client.Compress()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `str` |  |

#### Example: Create

```python
compress = client.Compress().create({
    "file": "example_file",  # str
})
```


### FillForm

Create an instance: `fill_form = client.FillForm()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `fields` | `str` |  |
| `file` | `str` |  |

#### Example: Create

```python
fill_form = client.FillForm().create({
    "fields": "example_fields",  # str
    "file": "example_file",  # str
})
```


### Health

Create an instance: `health = client.Health()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```python
health = client.Health().load()
```


### Html

Create an instance: `html = client.Html()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `filename` | `str` |  |
| `html` | `str` |  |

#### Example: Create

```python
html = client.Html().create({
    "html": "example_html",  # str
})
```


### Inbound

Create an instance: `inbound = client.Inbound()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```python
inbound = client.Inbound().create({
})
```


### Markdown

Create an instance: `markdown = client.Markdown()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `filename` | `str` |  |
| `markdown` | `str` |  |

#### Example: Create

```python
markdown = client.Markdown().create({
    "markdown": "example_markdown",  # str
})
```


### Men

Create an instance: `men = client.Men()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```python
men = client.Men().load()
```


### Merge

Create an instance: `merge = client.Merge()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `files` | `list` |  |

#### Example: Create

```python
merge = client.Merge().create({
    "files": [],  # list
})
```


### Portal

Create an instance: `portal = client.Portal()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```python
portal = client.Portal().create({
})
```


### RootGet

Create an instance: `root_get = client.RootGet()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```python
root_get = client.RootGet().load()
```


### Signup

Create an instance: `signup = client.Signup()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `str` |  |

#### Example: Create

```python
signup = client.Signup().create({
    "email": "example_email",  # str
})
```


### Split

Create an instance: `split = client.Split()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `str` |  |

#### Example: Create

```python
split = client.Split().create({
    "file": "example_file",  # str
})
```


### Table

Create an instance: `table = client.Table()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `str` |  |

#### Example: Create

```python
table = client.Table().create({
    "file": "example_file",  # str
})
```


### Template

Create an instance: `template = client.Template()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `dict` |  |
| `filename` | `str` |  |
| `template` | `str` |  |

#### Example: Create

```python
template = client.Template().create({
    "data": {},  # dict
    "template": "example_template",  # str
})
```


### Text

Create an instance: `text = client.Text()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `str` |  |

#### Example: Create

```python
text = client.Text().create({
    "file": "example_file",  # str
})
```


### ToImage

Create an instance: `to_image = client.ToImage()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dpi` | `int` |  |
| `file` | `str` |  |

#### Example: Create

```python
to_image = client.ToImage().create({
    "file": "example_file",  # str
})
```


### Translate

Create an instance: `translate = client.Translate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `Any` |  |
| `output` | `str` |  |
| `target_language` | `str` |  |
| `text` | `Any` |  |

#### Example: Create

```python
translate = client.Translate().create({
    "target_language": "example_target_language",  # str
})
```


### Watermark

Create an instance: `watermark = client.Watermark()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `str` |  |
| `text` | `str` |  |

#### Example: Create

```python
watermark = client.Watermark().create({
    "file": "example_file",  # str
})
```


### Webhook

Create an instance: `webhook = client.Webhook()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```python
webhook = client.Webhook().create({
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **RatelimitFeature**: Rate limiting
- **RetryFeature**: Retry
- **TestFeature**: Test transport
- **TimeoutFeature**: Timeout

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── inkspandocument_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── schema.py                    -- Generated option + entity specs
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`inkspandocument_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```python
health = client.Health()
health.load()

# health.data_get() now returns the health data from the last load
# health.match_get() returns the last match criteria
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
