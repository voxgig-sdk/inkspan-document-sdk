# InkspanDocument SDK configuration

module InkspanDocumentConfig
  # Return the process-wide config, built once on first use. The SDK reads
  # the config on every request and never writes to it, so one instance is
  # shared by every client rather than rebuilt per client.
  #
  # The returned hash is shared: treat it as read-only. Callers that need to
  # mutate should use make_config, which always returns a fresh copy.
  def self.shared_config
    @shared_config ||= make_config
  end


  # Build a fresh, fully materialised config hash. Every call rebuilds the
  # whole structure, so prefer shared_config unless you need a private copy
  # you intend to mutate.
  def self.make_config
    {
      "main" => {
        "name" => "InkspanDocument",
        "slug" => "inkspan-document",
        "version" => "0.0.1",
        "target" => "rb",
      },
      "feature" => {
        "ratelimit" => {
          "options" => {
            "active" => false,
            "burst" => 5,
            "rate" => 5,
          },
          "optspec" => {
            "now" => "`$FUNCTION`",
            "sleep" => "`$FUNCTION`",
          },
          "strict" => false,
          "transport" => "wrap",
        },
        "retry" => {
          "options" => {
            "active" => false,
            "factor" => 2,
            "maxDelay" => 2000,
            "minDelay" => 50,
            "retries" => 2,
            "statuses" => [
              408,
              425,
              429,
              500,
              502,
              503,
              504,
            ],
          },
          "optspec" => {
            "jitter" => "`$BOOLEAN`",
            "sleep" => "`$FUNCTION`",
          },
          "strict" => false,
          "transport" => "wrap",
        },
        "test" => {
          "options" => {
            "active" => false,
          },
          "optspec" => {
            "entity" => "`$MAP`",
            "net" => "`$MAP`",
          },
          "strict" => false,
          "transport" => "base",
        },
        "timeout" => {
          "options" => {
            "active" => false,
            "ms" => 30000,
          },
          "optspec" => {
            "clearTimer" => "`$FUNCTION`",
            "now" => "`$FUNCTION`",
            "setTimer" => "`$FUNCTION`",
          },
          "strict" => false,
          "transport" => "wrap",
        },
      },
      "options" => {
        "base" => "https://api.inkspan.dev",
        "headers" => {
          "content-type" => "application/json",
        },
        "entity" => {
          "checkout" => {},
          "compress" => {},
          "fill_form" => {},
          "health" => {},
          "html" => {},
          "inbound" => {},
          "markdown" => {},
          "men" => {},
          "merge" => {},
          "portal" => {},
          "root_get" => {},
          "signup" => {},
          "split" => {},
          "table" => {},
          "template" => {},
          "text" => {},
          "to_image" => {},
          "translate" => {},
          "watermark" => {},
          "webhook" => {},
        },
      },
      "entity" => {
        "checkout" => {
          "fields" => [],
          "name" => "checkout",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/billing/checkout",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "billing",
                    },
                    {
                      "lit" => "checkout",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "billing",
                    "checkout",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                    "query" => [
                      {
                        "name" => "plan",
                        "orig" => "plan",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "reqd" => true,
                      },
                    ],
                  },
                  "select" => {
                    "exist" => [
                      "plan",
                    ],
                  },
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "compress" => {
          "fields" => [
            {
              "name" => "file",
              "title" => "File",
              "type" => "`$STRING`",
              "req" => true,
            },
          ],
          "name" => "compress",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/pdf/compress",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "pdf",
                    },
                    {
                      "lit" => "compress",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "pdf",
                    "compress",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "body" => {
                    "fields" => [
                      {
                        "binary" => true,
                        "media" => "application/octet-stream",
                        "name" => "file",
                      },
                    ],
                    "kind" => "multipart",
                    "media" => "multipart/form-data",
                  },
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "fill_form" => {
          "fields" => [
            {
              "name" => "fields",
              "title" => "Fields",
              "type" => "`$STRING`",
              "req" => true,
            },
            {
              "name" => "file",
              "title" => "File",
              "type" => "`$STRING`",
              "req" => true,
            },
          ],
          "name" => "fill_form",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/pdf/fill-form",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "pdf",
                    },
                    {
                      "lit" => "fill-form",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "pdf",
                    "fill-form",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "body" => {
                    "fields" => [
                      {
                        "name" => "fields",
                      },
                      {
                        "binary" => true,
                        "media" => "application/octet-stream",
                        "name" => "file",
                      },
                    ],
                    "kind" => "multipart",
                    "media" => "multipart/form-data",
                  },
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "health" => {
          "fields" => [],
          "name" => "health",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/health",
                  "segments" => [
                    {
                      "lit" => "health",
                    },
                  ],
                  "parts" => [
                    "health",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {},
                  "select" => {},
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "html" => {
          "fields" => [
            {
              "name" => "filename",
              "title" => "Filename",
              "type" => "`$STRING`",
            },
            {
              "name" => "html",
              "title" => "Html",
              "type" => "`$STRING`",
              "req" => true,
            },
          ],
          "name" => "html",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/generate/html",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "generate",
                    },
                    {
                      "lit" => "html",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "generate",
                    "html",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "inbound" => {
          "fields" => [],
          "name" => "inbound",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/webhooks/inbound",
                  "segments" => [
                    {
                      "lit" => "webhooks",
                    },
                    {
                      "lit" => "inbound",
                    },
                  ],
                  "parts" => [
                    "webhooks",
                    "inbound",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "query" => [
                      {
                        "name" => "token",
                        "orig" => "token",
                        "type" => "`$STRING`",
                        "kind" => "query",
                        "example" => "",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "markdown" => {
          "fields" => [
            {
              "name" => "filename",
              "title" => "Filename",
              "type" => "`$STRING`",
            },
            {
              "name" => "markdown",
              "title" => "Markdown",
              "type" => "`$STRING`",
              "req" => true,
            },
          ],
          "name" => "markdown",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/generate/markdown",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "generate",
                    },
                    {
                      "lit" => "markdown",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "generate",
                    "markdown",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "men" => {
          "fields" => [],
          "name" => "men",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/v1/me",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "me",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "me",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "merge" => {
          "fields" => [
            {
              "name" => "files",
              "title" => "Files",
              "type" => "`$ARRAY`",
              "req" => true,
            },
          ],
          "name" => "merge",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/pdf/merge",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "pdf",
                    },
                    {
                      "lit" => "merge",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "pdf",
                    "merge",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "body" => {
                    "fields" => [
                      {
                        "binary" => true,
                        "list" => true,
                        "media" => "application/octet-stream",
                        "name" => "files",
                      },
                    ],
                    "kind" => "multipart",
                    "media" => "multipart/form-data",
                  },
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "portal" => {
          "fields" => [],
          "name" => "portal",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/billing/portal",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "billing",
                    },
                    {
                      "lit" => "portal",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "billing",
                    "portal",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "root_get" => {
          "fields" => [],
          "name" => "root_get",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "GET",
                  "orig" => "/",
                  "segments" => [],
                  "parts" => [],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {},
                  "select" => {},
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "signup" => {
          "fields" => [
            {
              "name" => "email",
              "title" => "Email",
              "type" => "`$STRING`",
              "req" => true,
            },
          ],
          "name" => "signup",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/signup",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "signup",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "signup",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {},
                  "select" => {},
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "split" => {
          "fields" => [
            {
              "name" => "file",
              "title" => "File",
              "type" => "`$STRING`",
              "req" => true,
            },
          ],
          "name" => "split",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/pdf/split",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "pdf",
                    },
                    {
                      "lit" => "split",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "pdf",
                    "split",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "body" => {
                    "fields" => [
                      {
                        "binary" => true,
                        "media" => "application/octet-stream",
                        "name" => "file",
                      },
                    ],
                    "kind" => "multipart",
                    "media" => "multipart/form-data",
                  },
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "table" => {
          "fields" => [
            {
              "name" => "file",
              "title" => "File",
              "type" => "`$STRING`",
              "req" => true,
            },
          ],
          "name" => "table",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/extract/tables",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "extract",
                    },
                    {
                      "lit" => "tables",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "extract",
                    "tables",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "body" => {
                    "fields" => [
                      {
                        "binary" => true,
                        "media" => "application/octet-stream",
                        "name" => "file",
                      },
                    ],
                    "kind" => "multipart",
                    "media" => "multipart/form-data",
                  },
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "template" => {
          "fields" => [
            {
              "name" => "data",
              "title" => "Data",
              "type" => "`$OBJECT`",
              "req" => true,
            },
            {
              "name" => "filename",
              "title" => "Filename",
              "type" => "`$STRING`",
            },
            {
              "name" => "template",
              "title" => "Template",
              "type" => "`$STRING`",
              "req" => true,
            },
          ],
          "name" => "template",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/generate/template",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "generate",
                    },
                    {
                      "lit" => "template",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "generate",
                    "template",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "text" => {
          "fields" => [
            {
              "name" => "file",
              "title" => "File",
              "type" => "`$STRING`",
              "req" => true,
            },
          ],
          "name" => "text",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/extract/text",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "extract",
                    },
                    {
                      "lit" => "text",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "extract",
                    "text",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "body" => {
                    "fields" => [
                      {
                        "binary" => true,
                        "media" => "application/octet-stream",
                        "name" => "file",
                      },
                    ],
                    "kind" => "multipart",
                    "media" => "multipart/form-data",
                  },
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "to_image" => {
          "fields" => [
            {
              "name" => "dpi",
              "title" => "Dpi",
              "type" => "`$INTEGER`",
            },
            {
              "name" => "file",
              "title" => "File",
              "type" => "`$STRING`",
              "req" => true,
            },
          ],
          "name" => "to_image",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/pdf/to-images",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "pdf",
                    },
                    {
                      "lit" => "to-images",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "pdf",
                    "to-images",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "body" => {
                    "fields" => [
                      {
                        "name" => "dpi",
                      },
                      {
                        "binary" => true,
                        "media" => "application/octet-stream",
                        "name" => "file",
                      },
                    ],
                    "kind" => "multipart",
                    "media" => "multipart/form-data",
                  },
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "translate" => {
          "fields" => [
            {
              "name" => "file",
              "title" => "File",
              "type" => "`$ANY`",
            },
            {
              "name" => "output",
              "title" => "Output",
              "type" => "`$STRING`",
            },
            {
              "name" => "target_language",
              "title" => "Target Language",
              "type" => "`$STRING`",
              "req" => true,
            },
            {
              "name" => "text",
              "title" => "Text",
              "type" => "`$ANY`",
            },
          ],
          "name" => "translate",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/translate",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "translate",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "translate",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "body" => {
                    "fields" => [
                      {
                        "name" => "file",
                      },
                      {
                        "name" => "output",
                      },
                      {
                        "name" => "target_language",
                      },
                      {
                        "name" => "text",
                      },
                    ],
                    "kind" => "multipart",
                    "media" => "multipart/form-data",
                  },
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "watermark" => {
          "fields" => [
            {
              "name" => "file",
              "title" => "File",
              "type" => "`$STRING`",
              "req" => true,
            },
            {
              "name" => "text",
              "title" => "Text",
              "type" => "`$STRING`",
            },
          ],
          "name" => "watermark",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/pdf/watermark",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "pdf",
                    },
                    {
                      "lit" => "watermark",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "pdf",
                    "watermark",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {
                    "header" => [
                      {
                        "name" => "x_api_key",
                        "orig" => "x-api-key",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                      {
                        "name" => "x_rapidapi_proxy_secret",
                        "orig" => "x-rapidapi-proxy-secret",
                        "type" => "`$ANY`",
                        "kind" => "header",
                      },
                    ],
                  },
                  "select" => {},
                  "body" => {
                    "fields" => [
                      {
                        "binary" => true,
                        "media" => "application/octet-stream",
                        "name" => "file",
                      },
                      {
                        "name" => "text",
                      },
                    ],
                    "kind" => "multipart",
                    "media" => "multipart/form-data",
                  },
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "webhook" => {
          "fields" => [],
          "name" => "webhook",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "kind" => "http",
                  "method" => "POST",
                  "orig" => "/v1/billing/webhook",
                  "segments" => [
                    {
                      "lit" => "v1",
                    },
                    {
                      "lit" => "billing",
                    },
                    {
                      "lit" => "webhook",
                    },
                  ],
                  "parts" => [
                    "v1",
                    "billing",
                    "webhook",
                  ],
                  "rename" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "args" => {},
                  "select" => {},
                  "response" => {
                    "kind" => "json",
                    "media" => "application/json",
                  },
                },
              ],
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
      },
    }
  end


  def self.make_feature(name)
    require_relative 'features'
    InkspanDocumentFeatures.make_feature(name)
  end
end
