# InkspanDocument PHP SDK



The PHP SDK for the InkspanDocument API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->Checkout()` — with named operations (`load`/`create`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`, see [Tags](https://github.com/voxgig-sdk/inkspan-document-sdk/tags)), or
from a clone as a Composer path repository:

```bash
git clone https://github.com/voxgig-sdk/inkspan-document-sdk
composer config repositories.inkspan-document-sdk path ./inkspan-document-sdk/php
composer require voxgig-sdk/inkspan-document-sdk:@dev
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'inkspandocument_sdk.php';

$client = new InkspanDocumentSDK();
```

### 4. Create, update, and remove

```php
// create() returns the ENTITY — call data_get() for the created Checkout record.
$created = $client->Checkout()->create(["plan" => "example_plan"]);

```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $health = $client->Health()->load();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required:

```php
$client = InkspanDocumentSDK::test();

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$health = $client->Health()->load();
print_r($health->data_get());
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new InkspanDocumentSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
INKSPAN_DOCUMENT_TEST_LIVE=TRUE
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### InkspanDocumentSDK

```php
require_once 'inkspandocument_sdk.php';
$client = new InkspanDocumentSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = InkspanDocumentSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### InkspanDocumentSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Checkout` | `($data): CheckoutEntity` | Create a Checkout entity instance. |
| `Compress` | `($data): CompressEntity` | Create a Compress entity instance. |
| `FillForm` | `($data): FillFormEntity` | Create a FillForm entity instance. |
| `Health` | `($data): HealthEntity` | Create a Health entity instance. |
| `Html` | `($data): HtmlEntity` | Create a Html entity instance. |
| `Inbound` | `($data): InboundEntity` | Create an Inbound entity instance. |
| `Markdown` | `($data): MarkdownEntity` | Create a Markdown entity instance. |
| `Men` | `($data): MenEntity` | Create a Men entity instance. |
| `Merge` | `($data): MergeEntity` | Create a Merge entity instance. |
| `Portal` | `($data): PortalEntity` | Create a Portal entity instance. |
| `RootGet` | `($data): RootGetEntity` | Create a RootGet entity instance. |
| `Signup` | `($data): SignupEntity` | Create a Signup entity instance. |
| `Split` | `($data): SplitEntity` | Create a Split entity instance. |
| `Table` | `($data): TableEntity` | Create a Table entity instance. |
| `Template` | `($data): TemplateEntity` | Create a Template entity instance. |
| `Text` | `($data): TextEntity` | Create a Text entity instance. |
| `ToImage` | `($data): ToImageEntity` | Create a ToImage entity instance. |
| `Translate` | `($data): TranslateEntity` | Create a Translate entity instance. |
| `Watermark` | `($data): WatermarkEntity` | Create a Watermark entity instance. |
| `Webhook` | `($data): WebhookEntity` | Create a Webhook entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): mixed` | Load a single entity by match criteria, and return it. |
| `create` | `($reqdata, $ctrl): mixed` | Create a new entity, and return it. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the entity, and `list` an `array` of entities, one
per record; an entity's `data_get()` reads its record (an `array`). They
throw on error, so wrap calls in `try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$checkout = $client->Checkout();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```php
$checkout = $client->Checkout()->create([
    "plan" => null, // string
]);
```


### Compress

Create an instance: `$compress = $client->Compress();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```php
$compress = $client->Compress()->create([
    "file" => null, // string
]);
```


### FillForm

Create an instance: `$fill_form = $client->FillForm();`

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

```php
$fill_form = $client->FillForm()->create([
    "fields" => null, // string
    "file" => null, // string
]);
```


### Health

Create an instance: `$health = $client->Health();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Health record (throws on error).
$health = $client->Health()->load();
```


### Html

Create an instance: `$html = $client->Html();`

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

```php
$html = $client->Html()->create([
    "html" => null, // string
]);
```


### Inbound

Create an instance: `$inbound = $client->Inbound();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```php
$inbound = $client->Inbound()->create([
]);
```


### Markdown

Create an instance: `$markdown = $client->Markdown();`

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

```php
$markdown = $client->Markdown()->create([
    "markdown" => null, // string
]);
```


### Men

Create an instance: `$men = $client->Men();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Men record (throws on error).
$men = $client->Men()->load();
```


### Merge

Create an instance: `$merge = $client->Merge();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `files` | `array` |  |

#### Example: Create

```php
$merge = $client->Merge()->create([
    "files" => null, // array
]);
```


### Portal

Create an instance: `$portal = $client->Portal();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```php
$portal = $client->Portal()->create([
]);
```


### RootGet

Create an instance: `$root_get = $client->RootGet();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the RootGet record (throws on error).
$root_get = $client->RootGet()->load();
```


### Signup

Create an instance: `$signup = $client->Signup();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` |  |

#### Example: Create

```php
$signup = $client->Signup()->create([
    "email" => null, // string
]);
```


### Split

Create an instance: `$split = $client->Split();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```php
$split = $client->Split()->create([
    "file" => null, // string
]);
```


### Table

Create an instance: `$table = $client->Table();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```php
$table = $client->Table()->create([
    "file" => null, // string
]);
```


### Template

Create an instance: `$template = $client->Template();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `array` |  |
| `filename` | `string` |  |
| `template` | `string` |  |

#### Example: Create

```php
$template = $client->Template()->create([
    "data" => null, // array
    "template" => null, // string
]);
```


### Text

Create an instance: `$text = $client->Text();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```php
$text = $client->Text()->create([
    "file" => null, // string
]);
```


### ToImage

Create an instance: `$to_image = $client->ToImage();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dpi` | `int` |  |
| `file` | `string` |  |

#### Example: Create

```php
$to_image = $client->ToImage()->create([
    "file" => null, // string
]);
```


### Translate

Create an instance: `$translate = $client->Translate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `mixed` |  |
| `output` | `string` |  |
| `target_language` | `string` |  |
| `text` | `mixed` |  |

#### Example: Create

```php
$translate = $client->Translate()->create([
    "target_language" => null, // string
]);
```


### Watermark

Create an instance: `$watermark = $client->Watermark();`

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

```php
$watermark = $client->Watermark()->create([
    "file" => null, // string
]);
```


### Webhook

Create an instance: `$webhook = $client->Webhook();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```php
$webhook = $client->Webhook()->create([
]);
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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **RatelimitFeature**: Rate limiting
- **RetryFeature**: Retry
- **TestFeature**: Test transport
- **TimeoutFeature**: Timeout

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── inkspandocument_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── schema.php                     -- Generated option + entity specs
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`inkspandocument_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```php
$health = $client->Health();
$health->load();

// $health->data_get() now returns the health data from the last load
// $health->match_get() returns the last match criteria
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
