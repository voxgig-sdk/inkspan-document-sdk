# InkspanDocument Python SDK Reference

Complete API reference for the InkspanDocument Python SDK.


## InkspanDocumentSDK

### Constructor

```python
from inkspandocument_sdk import InkspanDocumentSDK

client = InkspanDocumentSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `InkspanDocumentSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = InkspanDocumentSDK.test()
```


### Instance Methods

#### `Checkout(data=None)`

Create a new `CheckoutEntity` instance. Pass `None` for no initial data.

#### `Compress(data=None)`

Create a new `CompressEntity` instance. Pass `None` for no initial data.

#### `FillForm(data=None)`

Create a new `FillFormEntity` instance. Pass `None` for no initial data.

#### `Health(data=None)`

Create a new `HealthEntity` instance. Pass `None` for no initial data.

#### `Html(data=None)`

Create a new `HtmlEntity` instance. Pass `None` for no initial data.

#### `Inbound(data=None)`

Create a new `InboundEntity` instance. Pass `None` for no initial data.

#### `Markdown(data=None)`

Create a new `MarkdownEntity` instance. Pass `None` for no initial data.

#### `Men(data=None)`

Create a new `MenEntity` instance. Pass `None` for no initial data.

#### `Merge(data=None)`

Create a new `MergeEntity` instance. Pass `None` for no initial data.

#### `Portal(data=None)`

Create a new `PortalEntity` instance. Pass `None` for no initial data.

#### `RootGet(data=None)`

Create a new `RootGetEntity` instance. Pass `None` for no initial data.

#### `Signup(data=None)`

Create a new `SignupEntity` instance. Pass `None` for no initial data.

#### `Split(data=None)`

Create a new `SplitEntity` instance. Pass `None` for no initial data.

#### `Table(data=None)`

Create a new `TableEntity` instance. Pass `None` for no initial data.

#### `Template(data=None)`

Create a new `TemplateEntity` instance. Pass `None` for no initial data.

#### `Text(data=None)`

Create a new `TextEntity` instance. Pass `None` for no initial data.

#### `ToImage(data=None)`

Create a new `ToImageEntity` instance. Pass `None` for no initial data.

#### `Translate(data=None)`

Create a new `TranslateEntity` instance. Pass `None` for no initial data.

#### `Watermark(data=None)`

Create a new `WatermarkEntity` instance. Pass `None` for no initial data.

#### `Webhook(data=None)`

Create a new `WebhookEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## CheckoutEntity

```python
checkout = client.Checkout()
```

### Operations

#### `create(reqdata, ctrl=None) -> CheckoutEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Checkout().create({
    "plan": "example_plan",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CheckoutEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CompressEntity

```python
compress = client.Compress()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `file` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> CompressEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Compress().create({
    "file": "example_file",  # str
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CompressEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FillFormEntity

```python
fill_form = client.FillForm()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fields` | `str` | Yes |  |
| `file` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> FillFormEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.FillForm().create({
    "fields": "example_fields",  # str
    "file": "example_file",  # str
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FillFormEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## HealthEntity

```python
health = client.Health()
```

### Operations

#### `load(reqmatch, ctrl=None) -> HealthEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Health().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HealthEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## HtmlEntity

```python
html = client.Html()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `filename` | `str` | No |  |
| `html` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> HtmlEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Html().create({
    "html": "example_html",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HtmlEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboundEntity

```python
inbound = client.Inbound()
```

### Operations

#### `create(reqdata, ctrl=None) -> InboundEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Inbound().create({
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InboundEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MarkdownEntity

```python
markdown = client.Markdown()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `filename` | `str` | No |  |
| `markdown` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> MarkdownEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Markdown().create({
    "markdown": "example_markdown",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MarkdownEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MenEntity

```python
men = client.Men()
```

### Operations

#### `load(reqmatch, ctrl=None) -> MenEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Men().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MenEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MergeEntity

```python
merge = client.Merge()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `files` | `list` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> MergeEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Merge().create({
    "files": [],  # list
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MergeEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PortalEntity

```python
portal = client.Portal()
```

### Operations

#### `create(reqdata, ctrl=None) -> PortalEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Portal().create({
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PortalEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RootGetEntity

```python
root_get = client.RootGet()
```

### Operations

#### `load(reqmatch, ctrl=None) -> RootGetEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.RootGet().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RootGetEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SignupEntity

```python
signup = client.Signup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> SignupEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Signup().create({
    "email": "example_email",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SignupEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SplitEntity

```python
split = client.Split()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `file` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> SplitEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Split().create({
    "file": "example_file",  # str
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SplitEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TableEntity

```python
table = client.Table()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `file` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> TableEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Table().create({
    "file": "example_file",  # str
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TableEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TemplateEntity

```python
template = client.Template()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `dict` | Yes |  |
| `filename` | `str` | No |  |
| `template` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> TemplateEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Template().create({
    "data": {},  # dict
    "template": "example_template",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TemplateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TextEntity

```python
text = client.Text()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `file` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> TextEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Text().create({
    "file": "example_file",  # str
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TextEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ToImageEntity

```python
to_image = client.ToImage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dpi` | `int` | No |  |
| `file` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> ToImageEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.ToImage().create({
    "file": "example_file",  # str
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ToImageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TranslateEntity

```python
translate = client.Translate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `file` | `Any` | No |  |
| `output` | `str` | No |  |
| `target_language` | `str` | Yes |  |
| `text` | `Any` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> TranslateEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Translate().create({
    "target_language": "example_target_language",  # str
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TranslateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## WatermarkEntity

```python
watermark = client.Watermark()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `file` | `str` | Yes |  |
| `text` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> WatermarkEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Watermark().create({
    "file": "example_file",  # str
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WatermarkEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## WebhookEntity

```python
webhook = client.Webhook()
```

### Operations

#### `create(reqdata, ctrl=None) -> WebhookEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Webhook().create({
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WebhookEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `ratelimit` | 0.0.1 | Rate limiting |
| `retry` | 0.0.1 | Retry |
| `test` | 0.0.1 | Test transport |
| `timeout` | 0.0.1 | Timeout |


Features are activated via the `feature` option:

```python
client = InkspanDocumentSDK({
    "feature": {
        "ratelimit": {"active": True},
        "retry": {"active": True},
        "test": {"active": True},
        "timeout": {"active": True},
    },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `ratelimit`

Rate limiting.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Retry.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

Test transport.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Timeout.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `now` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

