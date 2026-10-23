package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewRatelimitFeatureFunc func() Feature

var NewRetryFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewTimeoutFeatureFunc func() Feature

var NewCheckoutEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewCompressEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewFillFormEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewHealthEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewHtmlEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewInboundEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewMarkdownEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewMenEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewMergeEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewPortalEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewRootGetEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewSignupEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewSplitEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewTableEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewTemplateEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewTextEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewToImageEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewTranslateEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewWatermarkEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

var NewWebhookEntityFunc func(client *InkspanDocumentSDK, entopts map[string]any) InkspanDocumentEntity

