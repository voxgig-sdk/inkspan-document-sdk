package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "InkspanDocument",
			"slug": "inkspan-document",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"now": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://api.inkspan.dev",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"checkout": map[string]any{},
				"compress": map[string]any{},
				"fill_form": map[string]any{},
				"health": map[string]any{},
				"html": map[string]any{},
				"inbound": map[string]any{},
				"markdown": map[string]any{},
				"men": map[string]any{},
				"merge": map[string]any{},
				"portal": map[string]any{},
				"root_get": map[string]any{},
				"signup": map[string]any{},
				"split": map[string]any{},
				"table": map[string]any{},
				"template": map[string]any{},
				"text": map[string]any{},
				"to_image": map[string]any{},
				"translate": map[string]any{},
				"watermark": map[string]any{},
				"webhook": map[string]any{},
			},
		},
		"entity": map[string]any{
			"checkout": map[string]any{
				"fields": []any{},
				"name": "checkout",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/billing/checkout",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "billing",
									},
									map[string]any{
										"lit": "checkout",
									},
								},
								"parts": []any{
									"v1",
									"billing",
									"checkout",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
									"query": []any{
										map[string]any{
											"name": "plan",
											"orig": "plan",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"plan",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"compress": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "compress",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/pdf/compress",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "pdf",
									},
									map[string]any{
										"lit": "compress",
									},
								},
								"parts": []any{
									"v1",
									"pdf",
									"compress",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"body": map[string]any{
									"fields": []any{
										map[string]any{
											"binary": true,
											"media": "application/octet-stream",
											"name": "file",
										},
									},
									"kind": "multipart",
									"media": "multipart/form-data",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"fill_form": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "fields",
						"title": "Fields",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "fill_form",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/pdf/fill-form",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "pdf",
									},
									map[string]any{
										"lit": "fill-form",
									},
								},
								"parts": []any{
									"v1",
									"pdf",
									"fill-form",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"body": map[string]any{
									"fields": []any{
										map[string]any{
											"name": "fields",
										},
										map[string]any{
											"binary": true,
											"media": "application/octet-stream",
											"name": "file",
										},
									},
									"kind": "multipart",
									"media": "multipart/form-data",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"health": map[string]any{
				"fields": []any{},
				"name": "health",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/health",
								"segments": []any{
									map[string]any{
										"lit": "health",
									},
								},
								"parts": []any{
									"health",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"html": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "filename",
						"title": "Filename",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "html",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/generate/html",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "generate",
									},
									map[string]any{
										"lit": "html",
									},
								},
								"parts": []any{
									"v1",
									"generate",
									"html",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"inbound": map[string]any{
				"fields": []any{},
				"name": "inbound",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/webhooks/inbound",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
									map[string]any{
										"lit": "inbound",
									},
								},
								"parts": []any{
									"webhooks",
									"inbound",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "token",
											"orig": "token",
											"type": "`$STRING`",
											"kind": "query",
											"example": "",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"markdown": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "filename",
						"title": "Filename",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "markdown",
						"title": "Markdown",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "markdown",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/generate/markdown",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "generate",
									},
									map[string]any{
										"lit": "markdown",
									},
								},
								"parts": []any{
									"v1",
									"generate",
									"markdown",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"men": map[string]any{
				"fields": []any{},
				"name": "men",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v1/me",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "me",
									},
								},
								"parts": []any{
									"v1",
									"me",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"merge": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "files",
						"title": "Files",
						"type": "`$ARRAY`",
						"req": true,
					},
				},
				"name": "merge",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/pdf/merge",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "pdf",
									},
									map[string]any{
										"lit": "merge",
									},
								},
								"parts": []any{
									"v1",
									"pdf",
									"merge",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"body": map[string]any{
									"fields": []any{
										map[string]any{
											"binary": true,
											"list": true,
											"media": "application/octet-stream",
											"name": "files",
										},
									},
									"kind": "multipart",
									"media": "multipart/form-data",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"portal": map[string]any{
				"fields": []any{},
				"name": "portal",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/billing/portal",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "billing",
									},
									map[string]any{
										"lit": "portal",
									},
								},
								"parts": []any{
									"v1",
									"billing",
									"portal",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"root_get": map[string]any{
				"fields": []any{},
				"name": "root_get",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/",
								"segments": []any{},
								"parts": []any{},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"signup": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "signup",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/signup",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "signup",
									},
								},
								"parts": []any{
									"v1",
									"signup",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"split": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "split",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/pdf/split",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "pdf",
									},
									map[string]any{
										"lit": "split",
									},
								},
								"parts": []any{
									"v1",
									"pdf",
									"split",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"body": map[string]any{
									"fields": []any{
										map[string]any{
											"binary": true,
											"media": "application/octet-stream",
											"name": "file",
										},
									},
									"kind": "multipart",
									"media": "multipart/form-data",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"table": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "table",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/extract/tables",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "extract",
									},
									map[string]any{
										"lit": "tables",
									},
								},
								"parts": []any{
									"v1",
									"extract",
									"tables",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"body": map[string]any{
									"fields": []any{
										map[string]any{
											"binary": true,
											"media": "application/octet-stream",
											"name": "file",
										},
									},
									"kind": "multipart",
									"media": "multipart/form-data",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"template": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "filename",
						"title": "Filename",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template",
						"title": "Template",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "template",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/generate/template",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "generate",
									},
									map[string]any{
										"lit": "template",
									},
								},
								"parts": []any{
									"v1",
									"generate",
									"template",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"text": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "text",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/extract/text",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "extract",
									},
									map[string]any{
										"lit": "text",
									},
								},
								"parts": []any{
									"v1",
									"extract",
									"text",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"body": map[string]any{
									"fields": []any{
										map[string]any{
											"binary": true,
											"media": "application/octet-stream",
											"name": "file",
										},
									},
									"kind": "multipart",
									"media": "multipart/form-data",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"to_image": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "dpi",
						"title": "Dpi",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "to_image",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/pdf/to-images",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "pdf",
									},
									map[string]any{
										"lit": "to-images",
									},
								},
								"parts": []any{
									"v1",
									"pdf",
									"to-images",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"body": map[string]any{
									"fields": []any{
										map[string]any{
											"name": "dpi",
										},
										map[string]any{
											"binary": true,
											"media": "application/octet-stream",
											"name": "file",
										},
									},
									"kind": "multipart",
									"media": "multipart/form-data",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"translate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "output",
						"title": "Output",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "target_language",
						"title": "Target Language",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": "`$ANY`",
					},
				},
				"name": "translate",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/translate",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "translate",
									},
								},
								"parts": []any{
									"v1",
									"translate",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"body": map[string]any{
									"fields": []any{
										map[string]any{
											"name": "file",
										},
										map[string]any{
											"name": "output",
										},
										map[string]any{
											"name": "target_language",
										},
										map[string]any{
											"name": "text",
										},
									},
									"kind": "multipart",
									"media": "multipart/form-data",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"watermark": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": "`$STRING`",
					},
				},
				"name": "watermark",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/pdf/watermark",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "pdf",
									},
									map[string]any{
										"lit": "watermark",
									},
								},
								"parts": []any{
									"v1",
									"pdf",
									"watermark",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "x_api_key",
											"orig": "x-api-key",
											"type": "`$ANY`",
											"kind": "header",
										},
										map[string]any{
											"name": "x_rapidapi_proxy_secret",
											"orig": "x-rapidapi-proxy-secret",
											"type": "`$ANY`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"body": map[string]any{
									"fields": []any{
										map[string]any{
											"binary": true,
											"media": "application/octet-stream",
											"name": "file",
										},
										map[string]any{
											"name": "text",
										},
									},
									"kind": "multipart",
									"media": "multipart/form-data",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"webhook": map[string]any{
				"fields": []any{},
				"name": "webhook",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/v1/billing/webhook",
								"segments": []any{
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "billing",
									},
									map[string]any{
										"lit": "webhook",
									},
								},
								"parts": []any{
									"v1",
									"billing",
									"webhook",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
