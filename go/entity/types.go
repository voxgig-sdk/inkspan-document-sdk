// Typed models for the InkspanDocument SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/inkspan-document-sdk/go/core"
)

// Checkout is the typed data model for the checkout entity.
type Checkout struct {
}

// CheckoutCreateData is the typed request payload for Checkout.CreateTyped.
type CheckoutCreateData struct {
	Plan string `json:"plan"`
}

// Compress is the typed data model for the compress entity.
type Compress struct {
}

// CompressCreateData is the typed request payload for Compress.CreateTyped.
type CompressCreateData struct {
	File string `json:"file"`
}

// FillForm is the typed data model for the fill_form entity.
type FillForm struct {
}

// FillFormCreateData is the typed request payload for FillForm.CreateTyped.
type FillFormCreateData struct {
	Fields string `json:"fields"`
	File string `json:"file"`
}

// Health is the typed data model for the health entity.
type Health struct {
}

// HealthLoadMatch is the typed request payload for Health.LoadTyped.
type HealthLoadMatch struct {
}

// Html is the typed data model for the html entity.
type Html struct {
}

// HtmlCreateData is the typed request payload for Html.CreateTyped.
type HtmlCreateData struct {
	Filename *string `json:"filename,omitempty"`
	Html string `json:"html"`
}

// Inbound is the typed data model for the inbound entity.
type Inbound struct {
}

// InboundCreateData is the typed request payload for Inbound.CreateTyped.
type InboundCreateData struct {
	Token *string `json:"token,omitempty"`
}

// Markdown is the typed data model for the markdown entity.
type Markdown struct {
}

// MarkdownCreateData is the typed request payload for Markdown.CreateTyped.
type MarkdownCreateData struct {
	Filename *string `json:"filename,omitempty"`
	Markdown string `json:"markdown"`
}

// Men is the typed data model for the men entity.
type Men struct {
}

// MenLoadMatch is the typed request payload for Men.LoadTyped.
type MenLoadMatch struct {
}

// Merge is the typed data model for the merge entity.
type Merge struct {
}

// MergeCreateData is the typed request payload for Merge.CreateTyped.
type MergeCreateData struct {
	Files []any `json:"files"`
}

// Portal is the typed data model for the portal entity.
type Portal struct {
}

// PortalCreateData is the typed request payload for Portal.CreateTyped.
type PortalCreateData struct {
}

// RootGet is the typed data model for the root_get entity.
type RootGet struct {
}

// RootGetLoadMatch is the typed request payload for RootGet.LoadTyped.
type RootGetLoadMatch struct {
}

// Signup is the typed data model for the signup entity.
type Signup struct {
}

// SignupCreateData is the typed request payload for Signup.CreateTyped.
type SignupCreateData struct {
	Email string `json:"email"`
}

// Split is the typed data model for the split entity.
type Split struct {
}

// SplitCreateData is the typed request payload for Split.CreateTyped.
type SplitCreateData struct {
	File string `json:"file"`
}

// Table is the typed data model for the table entity.
type Table struct {
}

// TableCreateData is the typed request payload for Table.CreateTyped.
type TableCreateData struct {
	File string `json:"file"`
}

// Template is the typed data model for the template entity.
type Template struct {
}

// TemplateCreateData is the typed request payload for Template.CreateTyped.
type TemplateCreateData struct {
	Data map[string]any `json:"data"`
	Filename *string `json:"filename,omitempty"`
	Template string `json:"template"`
}

// Text is the typed data model for the text entity.
type Text struct {
}

// TextCreateData is the typed request payload for Text.CreateTyped.
type TextCreateData struct {
	File string `json:"file"`
}

// ToImage is the typed data model for the to_image entity.
type ToImage struct {
}

// ToImageCreateData is the typed request payload for ToImage.CreateTyped.
type ToImageCreateData struct {
	Dpi *int `json:"dpi,omitempty"`
	File string `json:"file"`
}

// Translate is the typed data model for the translate entity.
type Translate struct {
}

// TranslateCreateData is the typed request payload for Translate.CreateTyped.
type TranslateCreateData struct {
	File *any `json:"file,omitempty"`
	Output *string `json:"output,omitempty"`
	TargetLanguage string `json:"target_language"`
	Text *any `json:"text,omitempty"`
}

// Watermark is the typed data model for the watermark entity.
type Watermark struct {
}

// WatermarkCreateData is the typed request payload for Watermark.CreateTyped.
type WatermarkCreateData struct {
	File string `json:"file"`
	Text *string `json:"text,omitempty"`
}

// Webhook is the typed data model for the webhook entity.
type Webhook struct {
}

// WebhookCreateData is the typed request payload for Webhook.CreateTyped.
type WebhookCreateData struct {
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
