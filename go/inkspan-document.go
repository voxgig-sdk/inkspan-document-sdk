package voxgiginkspandocumentsdk

import (
	"github.com/voxgig-sdk/inkspan-document-sdk/go/core"
	"github.com/voxgig-sdk/inkspan-document-sdk/go/entity"
	"github.com/voxgig-sdk/inkspan-document-sdk/go/feature"
	_ "github.com/voxgig-sdk/inkspan-document-sdk/go/utility"
)

// Type aliases preserve external API.
type InkspanDocumentSDK = core.InkspanDocumentSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type InkspanDocumentEntity = core.InkspanDocumentEntity
type StreamItem = core.StreamItem
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type InkspanDocumentError = core.InkspanDocumentError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewRatelimitFeatureFunc = func() core.Feature {
		return feature.NewRatelimitFeature()
	}
	core.NewRetryFeatureFunc = func() core.Feature {
		return feature.NewRetryFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewTimeoutFeatureFunc = func() core.Feature {
		return feature.NewTimeoutFeature()
	}
	core.NewCheckoutEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewCheckoutEntity(client, entopts)
	}
	core.NewCompressEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewCompressEntity(client, entopts)
	}
	core.NewFillFormEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewFillFormEntity(client, entopts)
	}
	core.NewHealthEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewHealthEntity(client, entopts)
	}
	core.NewHtmlEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewHtmlEntity(client, entopts)
	}
	core.NewInboundEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewInboundEntity(client, entopts)
	}
	core.NewMarkdownEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewMarkdownEntity(client, entopts)
	}
	core.NewMenEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewMenEntity(client, entopts)
	}
	core.NewMergeEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewMergeEntity(client, entopts)
	}
	core.NewPortalEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewPortalEntity(client, entopts)
	}
	core.NewRootGetEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewRootGetEntity(client, entopts)
	}
	core.NewSignupEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewSignupEntity(client, entopts)
	}
	core.NewSplitEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewSplitEntity(client, entopts)
	}
	core.NewTableEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewTableEntity(client, entopts)
	}
	core.NewTemplateEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewTemplateEntity(client, entopts)
	}
	core.NewTextEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewTextEntity(client, entopts)
	}
	core.NewToImageEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewToImageEntity(client, entopts)
	}
	core.NewTranslateEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewTranslateEntity(client, entopts)
	}
	core.NewWatermarkEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewWatermarkEntity(client, entopts)
	}
	core.NewWebhookEntityFunc = func(client *core.InkspanDocumentSDK, entopts map[string]any) core.InkspanDocumentEntity {
		return entity.NewWebhookEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewInkspanDocumentSDK = core.NewInkspanDocumentSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewInkspanDocumentSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *InkspanDocumentSDK  { return NewInkspanDocumentSDK(nil) }
func Test() *InkspanDocumentSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewRatelimitFeature = feature.NewRatelimitFeature
var NewRetryFeature = feature.NewRetryFeature
var NewTestFeature = feature.NewTestFeature
var NewTimeoutFeature = feature.NewTimeoutFeature
