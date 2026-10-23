// Typed models for the InkspanDocument SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Checkout {
}

export interface CheckoutCreateData {
  plan: string
}

export interface Compress {
  file: string
}

export interface CompressCreateData {
  file: string
}

export interface FillForm {
  fields: string
  file: string
}

export interface FillFormCreateData {
  fields: string
  file: string
}

export interface Health {
}

export interface HealthLoadMatch {
}

export interface Html {
  filename?: string
  html: string
}

export interface HtmlCreateData {
  filename?: string
  html: string
}

export interface Inbound {
}

export interface InboundCreateData {
  token?: string
}

export interface Markdown {
  filename?: string
  markdown: string
}

export interface MarkdownCreateData {
  filename?: string
  markdown: string
}

export interface Men {
}

export interface MenLoadMatch {
}

export interface Merge {
  files: any[]
}

export interface MergeCreateData {
  files: any[]
}

export interface Portal {
}

export interface PortalCreateData {
}

export interface RootGet {
}

export interface RootGetLoadMatch {
}

export interface Signup {
  email: string
}

export interface SignupCreateData {
  email: string
}

export interface Split {
  file: string
}

export interface SplitCreateData {
  file: string
}

export interface Table {
  file: string
}

export interface TableCreateData {
  file: string
}

export interface Template {
  data: Record<string, any>
  filename?: string
  template: string
}

export interface TemplateCreateData {
  data: Record<string, any>
  filename?: string
  template: string
}

export interface Text {
  file: string
}

export interface TextCreateData {
  file: string
}

export interface ToImage {
  dpi?: number
  file: string
}

export interface ToImageCreateData {
  dpi?: number
  file: string
}

export interface Translate {
  file?: any
  output?: string
  target_language: string
  text?: any
}

export interface TranslateCreateData {
  file?: any
  output?: string
  target_language: string
  text?: any
}

export interface Watermark {
  file: string
  text?: string
}

export interface WatermarkCreateData {
  file: string
  text?: string
}

export interface Webhook {
}

export interface WebhookCreateData {
}

