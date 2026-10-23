# InkspanDocument TypeScript SDK



The TypeScript SDK for the InkspanDocument API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Checkout()` — each with a small set of operations (`load`, `create`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb` — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`, see [Tags](https://github.com/voxgig-sdk/inkspan-document-sdk/tags)), or from a
clone, which carries the compiled `dist/`:

```bash
git clone https://github.com/voxgig-sdk/inkspan-document-sdk
npm install ./inkspan-document-sdk/ts
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { InkspanDocumentSDK } from '@voxgig-sdk/inkspan-document-sdk'

const client = new InkspanDocumentSDK()
```

### 4. Create, update, and remove

```ts
// Create — returns the created Checkout ENTITY (.data() for the record)
const created = await client.Checkout().create({
  plan: 'example_plan',
})

```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const health = await client.Health().load()
  console.log(health.data())
} catch (err) {
  console.error('load failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
result envelope. Branch on `ok`; on failure `status` holds the HTTP status
(for error responses) and `err` holds the error:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (!result.ok) {
  console.error('request failed:', result.status, result.err)
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = InkspanDocumentSDK.test()

const health = await client.Health().load()
// health is the Health entity; .data() reads its mock record
console.log(health.data())
```

You can also use the instance method:

```ts
const client = new InkspanDocumentSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Health()

// First call runs the operation and stores its result
await entity.load()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new InkspanDocumentSDK({
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
INKSPAN_DOCUMENT_TEST_LIVE=TRUE
```

Then run:

```bash
cd ts && npm test
```

Live entity tests continue independent operations after errors and attempt
supported cleanup. Their final result reports failures and missing prerequisites
after the remaining work completes. The model and test inputs determine which
API operations the generated scenarios cover.


## Reference

### InkspanDocumentSDK

#### Constructor

```ts
new InkspanDocumentSDK(options?: {
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `Checkout(data?)` | `CheckoutEntity` | Create a Checkout entity instance. |
| `Compress(data?)` | `CompressEntity` | Create a Compress entity instance. |
| `FillForm(data?)` | `FillFormEntity` | Create a FillForm entity instance. |
| `Health(data?)` | `HealthEntity` | Create a Health entity instance. |
| `Html(data?)` | `HtmlEntity` | Create a Html entity instance. |
| `Inbound(data?)` | `InboundEntity` | Create an Inbound entity instance. |
| `Markdown(data?)` | `MarkdownEntity` | Create a Markdown entity instance. |
| `Men(data?)` | `MenEntity` | Create a Men entity instance. |
| `Merge(data?)` | `MergeEntity` | Create a Merge entity instance. |
| `Portal(data?)` | `PortalEntity` | Create a Portal entity instance. |
| `RootGet(data?)` | `RootGetEntity` | Create a RootGet entity instance. |
| `Signup(data?)` | `SignupEntity` | Create a Signup entity instance. |
| `Split(data?)` | `SplitEntity` | Create a Split entity instance. |
| `Table(data?)` | `TableEntity` | Create a Table entity instance. |
| `Template(data?)` | `TemplateEntity` | Create a Template entity instance. |
| `Text(data?)` | `TextEntity` | Create a Text entity instance. |
| `ToImage(data?)` | `ToImageEntity` | Create a ToImage entity instance. |
| `Translate(data?)` | `TranslateEntity` | Create a Translate entity instance. |
| `Watermark(data?)` | `WatermarkEntity` | Create a Watermark entity instance. |
| `Webhook(data?)` | `WebhookEntity` | Create a Webhook entity instance. |
| `tester(testopts?, sdkopts?)` | `InkspanDocumentSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `InkspanDocumentSDK.test(testopts?, sdkopts?)` | `InkspanDocumentSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria, and return it. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity, and return it. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): InkspanDocumentSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity itself — there is no result
envelope, and an entity's `data()` reads its record:

- `load` and `create` resolve to a single entity object.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### Checkout

| Field | Description |
| --- | --- |

Operations: create.

API path: `/v1/billing/checkout`

#### Compress

| Field | Description |
| --- | --- |
| `file` |  |

Operations: create.

API path: `/v1/pdf/compress`

#### FillForm

| Field | Description |
| --- | --- |
| `fields` |  |
| `file` |  |

Operations: create.

API path: `/v1/pdf/fill-form`

#### Health

| Field | Description |
| --- | --- |

Operations: load.

API path: `/health`

#### Html

| Field | Description |
| --- | --- |
| `filename` |  |
| `html` |  |

Operations: create.

API path: `/v1/generate/html`

#### Inbound

| Field | Description |
| --- | --- |

Operations: create.

API path: `/webhooks/inbound`

#### Markdown

| Field | Description |
| --- | --- |
| `filename` |  |
| `markdown` |  |

Operations: create.

API path: `/v1/generate/markdown`

#### Men

| Field | Description |
| --- | --- |

Operations: load.

API path: `/v1/me`

#### Merge

| Field | Description |
| --- | --- |
| `files` |  |

Operations: create.

API path: `/v1/pdf/merge`

#### Portal

| Field | Description |
| --- | --- |

Operations: create.

API path: `/v1/billing/portal`

#### RootGet

| Field | Description |
| --- | --- |

Operations: load.

API path: `/`

#### Signup

| Field | Description |
| --- | --- |
| `email` |  |

Operations: create.

API path: `/v1/signup`

#### Split

| Field | Description |
| --- | --- |
| `file` |  |

Operations: create.

API path: `/v1/pdf/split`

#### Table

| Field | Description |
| --- | --- |
| `file` |  |

Operations: create.

API path: `/v1/extract/tables`

#### Template

| Field | Description |
| --- | --- |
| `data` |  |
| `filename` |  |
| `template` |  |

Operations: create.

API path: `/v1/generate/template`

#### Text

| Field | Description |
| --- | --- |
| `file` |  |

Operations: create.

API path: `/v1/extract/text`

#### ToImage

| Field | Description |
| --- | --- |
| `dpi` |  |
| `file` |  |

Operations: create.

API path: `/v1/pdf/to-images`

#### Translate

| Field | Description |
| --- | --- |
| `file` |  |
| `output` |  |
| `target_language` |  |
| `text` |  |

Operations: create.

API path: `/v1/translate`

#### Watermark

| Field | Description |
| --- | --- |
| `file` |  |
| `text` |  |

Operations: create.

API path: `/v1/pdf/watermark`

#### Webhook

| Field | Description |
| --- | --- |

Operations: create.

API path: `/v1/billing/webhook`



## Entities


### Checkout

Create an instance: `const checkout = client.Checkout()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```ts
const checkout = await client.Checkout().create({
  plan: 'example_plan',
})
```


### Compress

Create an instance: `const compress = client.Compress()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```ts
const compress = await client.Compress().create({
  file: 'example_file',
})
```


### FillForm

Create an instance: `const fill_form = client.FillForm()`

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

```ts
const fill_form = await client.FillForm().create({
  fields: 'example_fields',
  file: 'example_file',
})
```


### Health

Create an instance: `const health = client.Health()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ts
const health = await client.Health().load()
```


### Html

Create an instance: `const html = client.Html()`

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

```ts
const html = await client.Html().create({
  html: 'example_html',
})
```


### Inbound

Create an instance: `const inbound = client.Inbound()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```ts
const inbound = await client.Inbound().create({
})
```


### Markdown

Create an instance: `const markdown = client.Markdown()`

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

```ts
const markdown = await client.Markdown().create({
  markdown: 'example_markdown',
})
```


### Men

Create an instance: `const men = client.Men()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ts
const men = await client.Men().load()
```


### Merge

Create an instance: `const merge = client.Merge()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `files` | `any[]` |  |

#### Example: Create

```ts
const merge = await client.Merge().create({
  files: [],
})
```


### Portal

Create an instance: `const portal = client.Portal()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```ts
const portal = await client.Portal().create({
})
```


### RootGet

Create an instance: `const root_get = client.RootGet()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ts
const root_get = await client.RootGet().load()
```


### Signup

Create an instance: `const signup = client.Signup()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` |  |

#### Example: Create

```ts
const signup = await client.Signup().create({
  email: 'example_email',
})
```


### Split

Create an instance: `const split = client.Split()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```ts
const split = await client.Split().create({
  file: 'example_file',
})
```


### Table

Create an instance: `const table = client.Table()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```ts
const table = await client.Table().create({
  file: 'example_file',
})
```


### Template

Create an instance: `const template = client.Template()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `Record<string, any>` |  |
| `filename` | `string` |  |
| `template` | `string` |  |

#### Example: Create

```ts
const template = await client.Template().create({
  data: {},
  template: 'example_template',
})
```


### Text

Create an instance: `const text = client.Text()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file` | `string` |  |

#### Example: Create

```ts
const text = await client.Text().create({
  file: 'example_file',
})
```


### ToImage

Create an instance: `const to_image = client.ToImage()`

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

```ts
const to_image = await client.ToImage().create({
  file: 'example_file',
})
```


### Translate

Create an instance: `const translate = client.Translate()`

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

```ts
const translate = await client.Translate().create({
  target_language: 'example_target_language',
})
```


### Watermark

Create an instance: `const watermark = client.Watermark()`

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

```ts
const watermark = await client.Watermark().create({
  file: 'example_file',
})
```


### Webhook

Create an instance: `const webhook = client.Webhook()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```ts
const webhook = await client.Webhook().create({
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **RatelimitFeature**: Rate limiting
- **RetryFeature**: Retry
- **TestFeature**: Test transport
- **TimeoutFeature**: Timeout

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
inkspan-document/
├── src/
│   ├── InkspanDocumentSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { InkspanDocumentSDK } from '@voxgig-sdk/inkspan-document-sdk'
```

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const health = client.Health()
await health.load()

// health.data() now returns the health data from the last `load`
// health.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
