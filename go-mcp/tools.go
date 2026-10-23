package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	sdk "github.com/voxgig-sdk/inkspan-document-sdk/go"
)

// LoadArgs is what an agent sends to inkspan-document_load.
type LoadArgs struct {
	Entity string         `json:"entity" jsonschema:"one of: health | men | root_get"`
	Query  map[string]any `json:"query" jsonschema:"match map naming the record, such as {\"id\":1}"`
}

func registerTools(server *mcp.Server, client *sdk.InkspanDocumentSDK) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "inkspan-document_load",
		Description: "Load one record from InkspanDocument. Args: entity, query (match map naming the record, such as {\"id\":1}). Returns the record as JSON.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
		InputSchema: entitySchema[LoadArgs]("health", "men", "root_get"),
	}, func(ctx context.Context, req *mcp.CallToolRequest, args LoadArgs) (*mcp.CallToolResult, any, error) {
		return runOp(ctx, client, "load", args.Entity, args.Query)
	})
}

// entitySchema is the schema inferred from In, its entity limited to the
// entities the tool serves.
func entitySchema[In any](names ...string) *jsonschema.Schema {
	schema, err := jsonschema.For[In](nil)
	if err != nil {
		panic(err)
	}
	enum := make([]any, len(names))
	for i, name := range names {
		enum[i] = name
	}
	schema.Properties["entity"].Enum = enum
	return schema
}

func runOp(_ context.Context, client *sdk.InkspanDocumentSDK, op string, entity string, input map[string]any) (*mcp.CallToolResult, any, error) {
	ent, err := entityFor(client, entity)
	if err != nil {
		return toolError(err.Error())
	}

	var result any
	switch op {
	case "list":
		result, err = ent.List(input, nil)
	case "load":
		result, err = ent.Load(input, nil)
	case "create":
		result, err = ent.Create(input, nil)
	case "update":
		result, err = ent.Update(input, nil)
	case "patch":
		result, err = ent.Patch(input, nil)
	case "remove":
		result, err = ent.Remove(input, nil)
	default:
		return toolError(fmt.Sprintf("unknown op %q", op))
	}
	if err != nil {
		return toolError(err.Error())
	}

	// SDK returns *Entity wrappers; unwrap each via .Data() to get a
	// plain map[string]any (or []any of maps for list) suitable for
	// JSON marshalling.
	data := extractData(result)
	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return toolError(fmt.Sprintf("marshal: %v", err))
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(body)},
		},
	}, data, nil
}

// entityFor dispatches on the lowercase entity name. The generator
// emits one `case "<name>":` per entity defined in the SDK model.
func entityFor(client *sdk.InkspanDocumentSDK, name string) (sdk.InkspanDocumentEntity, error) {
	switch strings.ToLower(name) {
	case "checkout":
		return client.Checkout(nil), nil
	case "compress":
		return client.Compress(nil), nil
	case "fill_form":
		return client.FillForm(nil), nil
	case "health":
		return client.Health(nil), nil
	case "html":
		return client.Html(nil), nil
	case "inbound":
		return client.Inbound(nil), nil
	case "markdown":
		return client.Markdown(nil), nil
	case "men":
		return client.Men(nil), nil
	case "merge":
		return client.Merge(nil), nil
	case "portal":
		return client.Portal(nil), nil
	case "root_get":
		return client.RootGet(nil), nil
	case "signup":
		return client.Signup(nil), nil
	case "split":
		return client.Split(nil), nil
	case "table":
		return client.Table(nil), nil
	case "template":
		return client.Template(nil), nil
	case "text":
		return client.Text(nil), nil
	case "to_image":
		return client.ToImage(nil), nil
	case "translate":
		return client.Translate(nil), nil
	case "watermark":
		return client.Watermark(nil), nil
	case "webhook":
		return client.Webhook(nil), nil
	}
	return nil, fmt.Errorf("unknown entity %q", name)
}

func extractData(x any) any {
	switch v := x.(type) {
	case sdk.Entity:
		return extractData(v.Data())
	case []any:
		out := make([]any, len(v))
		for i, e := range v {
			out[i] = extractData(e)
		}
		return out
	case map[string]any:
		out := make(map[string]any, len(v))
		for k, vv := range v {
			out[k] = extractData(vv)
		}
		return out
	}
	return x
}

func toolError(msg string) (*mcp.CallToolResult, any, error) {
	return &mcp.CallToolResult{
		IsError: true,
		Content: []mcp.Content{
			&mcp.TextContent{Text: msg},
		},
	}, nil, nil
}

// hint is an MCP annotation that defaults to true unless stated.
func hint(b bool) *bool {
	return &b
}
