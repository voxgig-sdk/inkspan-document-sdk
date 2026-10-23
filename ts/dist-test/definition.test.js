"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const node_test_1 = require("node:test");
const __1 = require("..");
const definition_runner_1 = require("./definition-runner");
const utility_1 = require("./utility");
// Generated from the API definition, not from the model this SDK was built
// from: the route, the declared query parameters, the credential the security
// scheme names, and the definition's own response example.
const PLAN = [
    {
        "entity": "checkout",
        "accessor": "Checkout",
        "op": "create",
        "method": "POST",
        "path": "/v1/billing/checkout",
        "args": [],
        "select": {
            "plan": "v1"
        },
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [
            "plan"
        ],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "compress",
        "accessor": "Compress",
        "op": "create",
        "method": "POST",
        "path": "/v1/pdf/compress",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "fill_form",
        "accessor": "FillForm",
        "op": "create",
        "method": "POST",
        "path": "/v1/pdf/fill-form",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "health",
        "accessor": "Health",
        "op": "load",
        "method": "GET",
        "path": "/health",
        "args": [],
        "select": {},
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "html",
        "accessor": "Html",
        "op": "create",
        "method": "POST",
        "path": "/v1/generate/html",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "inbound",
        "accessor": "Inbound",
        "op": "create",
        "method": "POST",
        "path": "/webhooks/inbound",
        "args": [],
        "select": {
            "token": ""
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [
            "token"
        ],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "markdown",
        "accessor": "Markdown",
        "op": "create",
        "method": "POST",
        "path": "/v1/generate/markdown",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "men",
        "accessor": "Men",
        "op": "load",
        "method": "GET",
        "path": "/v1/me",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "merge",
        "accessor": "Merge",
        "op": "create",
        "method": "POST",
        "path": "/v1/pdf/merge",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "portal",
        "accessor": "Portal",
        "op": "create",
        "method": "POST",
        "path": "/v1/billing/portal",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "root_get",
        "accessor": "RootGet",
        "op": "load",
        "method": "GET",
        "path": "/",
        "args": [],
        "select": {},
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "signup",
        "accessor": "Signup",
        "op": "create",
        "method": "POST",
        "path": "/v1/signup",
        "args": [],
        "select": {},
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "split",
        "accessor": "Split",
        "op": "create",
        "method": "POST",
        "path": "/v1/pdf/split",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "table",
        "accessor": "Table",
        "op": "create",
        "method": "POST",
        "path": "/v1/extract/tables",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "template",
        "accessor": "Template",
        "op": "create",
        "method": "POST",
        "path": "/v1/generate/template",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "text",
        "accessor": "Text",
        "op": "create",
        "method": "POST",
        "path": "/v1/extract/text",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "to_image",
        "accessor": "ToImage",
        "op": "create",
        "method": "POST",
        "path": "/v1/pdf/to-images",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "translate",
        "accessor": "Translate",
        "op": "create",
        "method": "POST",
        "path": "/v1/translate",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "watermark",
        "accessor": "Watermark",
        "op": "create",
        "method": "POST",
        "path": "/v1/pdf/watermark",
        "args": [],
        "select": {},
        "headers": [
            {
                "name": "x_api_key",
                "wire": "x-api-key",
                "value": "h1"
            },
            {
                "name": "x_rapidapi_proxy_secret",
                "wire": "x-rapidapi-proxy-secret",
                "value": "h2"
            }
        ],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    },
    {
        "entity": "webhook",
        "accessor": "Webhook",
        "op": "create",
        "method": "POST",
        "path": "/v1/billing/webhook",
        "args": [],
        "select": {},
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json"
        ],
        "query": [],
        "queryArgs": [],
        "auth": null,
        "status": 200,
        "sample": null,
        "idField": "id"
    }
];
(0, node_test_1.describe)('definition', () => {
    for (const point of PLAN) {
        (0, node_test_1.test)(point.entity + '.' + point.op + ' ' + point.method + ' ' + point.path, async (t) => {
            const control = (0, utility_1.isControlSkipped)('entityOp', point.entity + '.' + point.op, 'definition');
            if (control.skip) {
                t.skip(control.reason || 'skipped via sdk-test-control.json');
                return;
            }
            await (0, definition_runner_1.runDefinitionPoint)(__1.SDK, point);
        });
    }
});
//# sourceMappingURL=definition.test.js.map