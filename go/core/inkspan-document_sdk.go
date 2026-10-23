package core

import (
	"encoding/json"
	"fmt"
	"strings"

	vs "github.com/voxgig-sdk/inkspan-document-sdk/go/utility/struct"
)

type InkspanDocumentSDK struct {
	Mode     string
	options  map[string]any
	utility  *Utility
	Features []Feature
	rootctx  *Context
}

func NewInkspanDocumentSDK(options map[string]any) *InkspanDocumentSDK {
	sdk := &InkspanDocumentSDK{
		Mode:     "live",
		Features: []Feature{},
	}

	sdk.utility = NewUtility()

	config := SharedConfig()

	sdk.rootctx = sdk.utility.MakeContext(map[string]any{
		"client":  sdk,
		"utility": sdk.utility,
		"config":  config,
		"options": options,
		"shared":  map[string]any{},
	}, nil)

	sdk.options = sdk.utility.MakeOptions(sdk.rootctx)

	if vs.GetPath(sdk.options, []any{"feature", "test", "active"}) == true {
		sdk.Mode = "test"
	}

	sdk.rootctx.Options = sdk.options

	// Add features in the resolved order (MakeOptions puts an explicit array
	// order first, else defaults to test-first). Ordering matters: the `test`
	// feature installs the base mock transport and the transport features
	// (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
	// must be added before them to sit at the base of the chain.
	featureOpts := ToMapAny(vs.GetProp(sdk.options, "feature"))
	if featureOpts != nil {
		if fo, ok := vs.GetPath(sdk.options, []any{"__derived__", "featureorder"}).([]any); ok {
			for _, n := range fo {
				fname, _ := n.(string)
				fopts := ToMapAny(featureOpts[fname])
				if fopts != nil {
					if active, ok := fopts["active"]; ok {
						if ab, ok := active.(bool); ok && ab {
							sdk.utility.FeatureAdd(sdk.rootctx, makeFeature(fname))
						}
					}
				}
			}
		}
	}

	// Add extension features.
	if extend := vs.GetProp(sdk.options, "extend"); extend != nil {
		if extList, ok := extend.([]any); ok {
			for _, f := range extList {
				if feat, ok := f.(Feature); ok {
					sdk.utility.FeatureAdd(sdk.rootctx, feat)
				}
			}
		}
	}

	// Initialize features.
	for _, f := range sdk.Features {
		sdk.utility.FeatureInit(sdk.rootctx, f)
	}

	sdk.utility.FeatureHook(sdk.rootctx, "PostConstruct")

	return sdk
}

// The client holds the credential in its options, so a print or a JSON dump
// carries the name alone. Value receivers: a dereferenced client prints the
// same way.
func (sdk InkspanDocumentSDK) String() string {
	return "InkspanDocument " + vs.Jsonify(map[string]any{"name": "InkspanDocument"},
		map[string]any{"indent": 0})
}

func (sdk InkspanDocumentSDK) GoString() string {
	return sdk.String()
}

func (sdk InkspanDocumentSDK) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{"name": "InkspanDocument"})
}

func (sdk *InkspanDocumentSDK) OptionsMap() map[string]any {
	out := vs.Clone(sdk.options)
	if om, ok := out.(map[string]any); ok {
		return om
	}
	return map[string]any{}
}

func (sdk *InkspanDocumentSDK) GetUtility() *Utility {
	return CopyUtility(sdk.utility)
}

func (sdk *InkspanDocumentSDK) GetRootCtx() *Context {
	return sdk.rootctx
}

func (sdk *InkspanDocumentSDK) Prepare(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "prepare",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	options := sdk.options

	path, _ := vs.GetProp(fetchargs, "path").(string)
	method, _ := vs.GetProp(fetchargs, "method").(string)
	if method == "" {
		method = "GET"
	}
	method = strings.ToUpper(method)

	allowMethodVal := vs.GetPath(options, []any{"allow", "method"})
	if !Allowed(allowMethodVal, method) {
		allowMethod, _ := allowMethodVal.(string)
		return nil, ctx.MakeError("spec_method_allow",
			"Method \""+method+"\" not allowed by SDK option allow.method value: \""+allowMethod+"\"")
	}

	params := ToMapAny(vs.GetProp(fetchargs, "params"))
	if params == nil {
		params = map[string]any{}
	}
	query := ToMapAny(vs.GetProp(fetchargs, "query"))
	if query == nil {
		query = map[string]any{}
	}

	headers := utility.PrepareHeaders(ctx)

	base, _ := vs.GetProp(options, "base").(string)
	prefix, _ := vs.GetProp(options, "prefix").(string)
	suffix, _ := vs.GetProp(options, "suffix").(string)

	ctx.Spec = NewSpec(map[string]any{
		"base":    base,
		"prefix":  prefix,
		"suffix":  suffix,
		"path":    path,
		"method":  method,
		"params":  params,
		"query":   query,
		"headers": headers,
		"body":    vs.GetProp(fetchargs, "body"),
		"step":    "start",
	})

	// Merge user-provided headers.
	if uh := vs.GetProp(fetchargs, "headers"); uh != nil {
		if uhm, ok := uh.(map[string]any); ok {
			for k, v := range uhm {
				ctx.Spec.Headers[k] = v
			}
		}
	}

	_, err := utility.PrepareAuth(ctx)
	if err != nil {
		return nil, err
	}

	return utility.MakeFetchDef(ctx)
}

// Raw endpoint access is operator-controllable, like every entity op.
// Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
// either one reaches the same endpoint.
func (sdk *InkspanDocumentSDK) Direct(fetchargs map[string]any) (map[string]any, error) {
	if !sdk.opAllowed("direct") {
		return sdk.opDenied("direct"), nil
	}

	return sdk.rawRequest(fetchargs)
}

// Is this raw-access op permitted by the SDK's allow.op option?
func (sdk *InkspanDocumentSDK) opAllowed(op string) bool {
	return Allowed(vs.GetPath(sdk.options, []any{"allow", "op"}), op)
}

func (sdk *InkspanDocumentSDK) opDenied(op string) map[string]any {
	allowOp, _ := vs.GetPath(sdk.options, []any{"allow", "op"}).(string)
	return map[string]any{
		"ok": false,
		"err": fmt.Errorf("InkspanDocumentSDK: %s: operation not allowed by"+
			" SDK option allow.op value: \"%s\"", op, allowOp),
	}
}

// Ungated request path shared by Direct and Graphql, each of which checks
// its own allow.op token first. Unexported, rather than a flag on fetchargs:
// a caller-supplied marker would let anyone opt straight back out of the
// gate by passing it.
func (sdk *InkspanDocumentSDK) rawRequest(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	fetchdef, err := sdk.Prepare(fetchargs)
	if err != nil {
		return map[string]any{"ok": false, "err": sdk.cleanErr(sdk.rootctx, err)}, nil
	}

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "direct",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	url, _ := fetchdef["url"].(string)
	fetched, fetchErr := utility.Fetcher(ctx, url, fetchdef)

	if fetchErr != nil {
		return map[string]any{"ok": false, "err": sdk.cleanErr(ctx, fetchErr)}, nil
	}

	if fetched == nil {
		return map[string]any{
			"ok":  false,
			"err": ctx.MakeError("direct_no_response", "response: undefined"),
		}, nil
	}

	if fm, ok := fetched.(map[string]any); ok {
		status := ToInt(vs.GetProp(fm, "status"))
		headers := vs.GetProp(fm, "headers")

		// No-body responses (204, 304) and explicit zero content-length
		// must skip JSON parsing — calling json() on an empty body errors.
		var contentLength string
		if hm, ok := headers.(map[string]any); ok {
			if cl, ok := hm["content-length"]; ok {
				contentLength = fmt.Sprintf("%v", cl)
			}
		}
		noBody := status == 204 || status == 304 || contentLength == "0"

		var jsonData any
		var bodyErr error
		if !noBody {
			if jf := vs.GetProp(fm, "json"); jf != nil {
				if f, ok := jf.(func() any); ok {
					jsonData = f()
				}
			}
			if unreadable, _ := vs.GetProp(fm, "unreadable").(bool); unreadable {
				var failed error
				if status < 200 || status >= 300 {
					failed = ctx.MakeError("request_status",
						fmt.Sprintf("request: %d: %v", status, vs.GetProp(fm, "statusText")))
				}
				bodyErr = UnreadableBody(ctx, status, headers, vs.GetProp(fm, "body"),
					fetchdef["headers"], failed)
			}
		}

		out := map[string]any{
			"ok":      bodyErr == nil && status >= 200 && status < 300,
			"status":  status,
			"headers": headers,
			"data":    jsonData,
		}
		if bodyErr != nil {
			out["err"] = sdk.cleanErr(ctx, bodyErr)
		}
		return out, nil
	}

	return map[string]any{"ok": false, "err": ctx.MakeError("direct_invalid", "invalid response type")}, nil
}

// A raw request returns its error rather than passing it through MakeError.
func (sdk *InkspanDocumentSDK) cleanErr(ctx *Context, err error) error {
	if cleaned, ok := sdk.utility.Clean(ctx, err).(error); ok {
		return cleaned
	}
	return err
}

func (sdk *InkspanDocumentSDK) Graphql(
	query string, variables map[string]any, ctrl map[string]any,
) (map[string]any, error) {
	if !sdk.opAllowed("graphql") {
		return sdk.opDenied("graphql"), nil
	}

	if variables == nil {
		variables = map[string]any{}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	res, err := sdk.rawRequest(map[string]any{
		"method":  "POST",
		"headers": map[string]any{"content-type": "application/json"},
		"body":    map[string]any{"query": query, "variables": variables},
		"ctrl":    ctrl,
	})

	if err != nil {
		return res, err
	}

	// Errors are read BEFORE any status check: a GraphQL parse or validation
	// failure comes back as HTTP 400 carrying the standard { errors: [...] }
	// body, and the raw path represents a non-2xx as ok:false with no err —
	// so returning early on status would discard the server's own
	// diagnostics, which are the only useful part of that response.
	errors, _ := vs.GetPath(res, []any{"data", "errors"}).([]any)

	if 0 < len(errors) {
		msg, _ := vs.GetProp(errors[0], "message").(string)
		if msg == "" {
			msg = "graphql error"
		}
		res["ok"] = false
		res["err"] = fmt.Errorf("InkspanDocumentSDK: graphql: %s", msg)
		res["graphql"] = errors
	}

	return res, nil
}


// Checkout returns a Checkout entity bound to this client.
// Idiomatic usage: client.Checkout(nil).List(nil, nil) or
// client.Checkout(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Checkout(data map[string]any) InkspanDocumentEntity {
	return NewCheckoutEntityFunc(sdk, data)
}


// Compress returns a Compress entity bound to this client.
// Idiomatic usage: client.Compress(nil).List(nil, nil) or
// client.Compress(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Compress(data map[string]any) InkspanDocumentEntity {
	return NewCompressEntityFunc(sdk, data)
}


// FillForm returns a FillForm entity bound to this client.
// Idiomatic usage: client.FillForm(nil).List(nil, nil) or
// client.FillForm(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) FillForm(data map[string]any) InkspanDocumentEntity {
	return NewFillFormEntityFunc(sdk, data)
}


// Health returns a Health entity bound to this client.
// Idiomatic usage: client.Health(nil).List(nil, nil) or
// client.Health(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Health(data map[string]any) InkspanDocumentEntity {
	return NewHealthEntityFunc(sdk, data)
}


// Html returns a Html entity bound to this client.
// Idiomatic usage: client.Html(nil).List(nil, nil) or
// client.Html(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Html(data map[string]any) InkspanDocumentEntity {
	return NewHtmlEntityFunc(sdk, data)
}


// Inbound returns a Inbound entity bound to this client.
// Idiomatic usage: client.Inbound(nil).List(nil, nil) or
// client.Inbound(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Inbound(data map[string]any) InkspanDocumentEntity {
	return NewInboundEntityFunc(sdk, data)
}


// Markdown returns a Markdown entity bound to this client.
// Idiomatic usage: client.Markdown(nil).List(nil, nil) or
// client.Markdown(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Markdown(data map[string]any) InkspanDocumentEntity {
	return NewMarkdownEntityFunc(sdk, data)
}


// Men returns a Men entity bound to this client.
// Idiomatic usage: client.Men(nil).List(nil, nil) or
// client.Men(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Men(data map[string]any) InkspanDocumentEntity {
	return NewMenEntityFunc(sdk, data)
}


// Merge returns a Merge entity bound to this client.
// Idiomatic usage: client.Merge(nil).List(nil, nil) or
// client.Merge(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Merge(data map[string]any) InkspanDocumentEntity {
	return NewMergeEntityFunc(sdk, data)
}


// Portal returns a Portal entity bound to this client.
// Idiomatic usage: client.Portal(nil).List(nil, nil) or
// client.Portal(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Portal(data map[string]any) InkspanDocumentEntity {
	return NewPortalEntityFunc(sdk, data)
}


// RootGet returns a RootGet entity bound to this client.
// Idiomatic usage: client.RootGet(nil).List(nil, nil) or
// client.RootGet(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) RootGet(data map[string]any) InkspanDocumentEntity {
	return NewRootGetEntityFunc(sdk, data)
}


// Signup returns a Signup entity bound to this client.
// Idiomatic usage: client.Signup(nil).List(nil, nil) or
// client.Signup(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Signup(data map[string]any) InkspanDocumentEntity {
	return NewSignupEntityFunc(sdk, data)
}


// Split returns a Split entity bound to this client.
// Idiomatic usage: client.Split(nil).List(nil, nil) or
// client.Split(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Split(data map[string]any) InkspanDocumentEntity {
	return NewSplitEntityFunc(sdk, data)
}


// Table returns a Table entity bound to this client.
// Idiomatic usage: client.Table(nil).List(nil, nil) or
// client.Table(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Table(data map[string]any) InkspanDocumentEntity {
	return NewTableEntityFunc(sdk, data)
}


// Template returns a Template entity bound to this client.
// Idiomatic usage: client.Template(nil).List(nil, nil) or
// client.Template(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Template(data map[string]any) InkspanDocumentEntity {
	return NewTemplateEntityFunc(sdk, data)
}


// Text returns a Text entity bound to this client.
// Idiomatic usage: client.Text(nil).List(nil, nil) or
// client.Text(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Text(data map[string]any) InkspanDocumentEntity {
	return NewTextEntityFunc(sdk, data)
}


// ToImage returns a ToImage entity bound to this client.
// Idiomatic usage: client.ToImage(nil).List(nil, nil) or
// client.ToImage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) ToImage(data map[string]any) InkspanDocumentEntity {
	return NewToImageEntityFunc(sdk, data)
}


// Translate returns a Translate entity bound to this client.
// Idiomatic usage: client.Translate(nil).List(nil, nil) or
// client.Translate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Translate(data map[string]any) InkspanDocumentEntity {
	return NewTranslateEntityFunc(sdk, data)
}


// Watermark returns a Watermark entity bound to this client.
// Idiomatic usage: client.Watermark(nil).List(nil, nil) or
// client.Watermark(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Watermark(data map[string]any) InkspanDocumentEntity {
	return NewWatermarkEntityFunc(sdk, data)
}


// Webhook returns a Webhook entity bound to this client.
// Idiomatic usage: client.Webhook(nil).List(nil, nil) or
// client.Webhook(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *InkspanDocumentSDK) Webhook(data map[string]any) InkspanDocumentEntity {
	return NewWebhookEntityFunc(sdk, data)
}



func TestSDK(testopts map[string]any, sdkopts map[string]any) *InkspanDocumentSDK {
	if sdkopts == nil {
		sdkopts = map[string]any{}
	}
	sdkopts = vs.Clone(sdkopts).(map[string]any)

	if testopts == nil {
		testopts = map[string]any{}
	}
	testopts = vs.Clone(testopts).(map[string]any)
	testopts["active"] = true

	vs.SetPath(sdkopts, []any{"feature", "test"}, testopts)

	sdk := NewInkspanDocumentSDK(sdkopts)
	sdk.Mode = "test"

	return sdk
}
