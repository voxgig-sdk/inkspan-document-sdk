-- Typed models for the InkspanDocument SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
-- params (op.<name>.points[].g.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Checkout

---@class CheckoutCreateData
---@field plan string

---@class Compress
---@field file string

---@class CompressCreateData
---@field file string

---@class FillForm
---@field fields string
---@field file string

---@class FillFormCreateData
---@field fields string
---@field file string

---@class Health

---@class HealthLoadMatch

---@class Html
---@field filename? string
---@field html string

---@class HtmlCreateData
---@field filename? string
---@field html string

---@class Inbound

---@class InboundCreateData
---@field token? string

---@class Markdown
---@field filename? string
---@field markdown string

---@class MarkdownCreateData
---@field filename? string
---@field markdown string

---@class Men

---@class MenLoadMatch

---@class Merge
---@field files table

---@class MergeCreateData
---@field files table

---@class Portal

---@class PortalCreateData

---@class RootGet

---@class RootGetLoadMatch

---@class Signup
---@field email string

---@class SignupCreateData
---@field email string

---@class Split
---@field file string

---@class SplitCreateData
---@field file string

---@class Table
---@field file string

---@class TableCreateData
---@field file string

---@class Template
---@field data table
---@field filename? string
---@field template string

---@class TemplateCreateData
---@field data table
---@field filename? string
---@field template string

---@class Text
---@field file string

---@class TextCreateData
---@field file string

---@class ToImage
---@field dpi? number
---@field file string

---@class ToImageCreateData
---@field dpi? number
---@field file string

---@class Translate
---@field file? any
---@field output? string
---@field target_language string
---@field text? any

---@class TranslateCreateData
---@field file? any
---@field output? string
---@field target_language string
---@field text? any

---@class Watermark
---@field file string
---@field text? string

---@class WatermarkCreateData
---@field file string
---@field text? string

---@class Webhook

---@class WebhookCreateData

local M = {}

return M
