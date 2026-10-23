<?php
declare(strict_types=1);

// Typed models for the InkspanDocument SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Checkout entity data model. */
class Checkout
{
}

/** Request payload for Checkout#create. */
class CheckoutCreateData
{
    public string $plan;
}

/** Compress entity data model. */
class Compress
{
    public string $file;
}

/** Request payload for Compress#create. */
class CompressCreateData
{
    public string $file;
}

/** FillForm entity data model. */
class FillForm
{
    public string $fields;
    public string $file;
}

/** Request payload for FillForm#create. */
class FillFormCreateData
{
    public string $fields;
    public string $file;
}

/** Health entity data model. */
class Health
{
}

/** Request payload for Health#load. */
class HealthLoadMatch
{
}

/** Html entity data model. */
class Html
{
    public ?string $filename = null;
    public string $html;
}

/** Request payload for Html#create. */
class HtmlCreateData
{
    public ?string $filename = null;
    public string $html;
}

/** Inbound entity data model. */
class Inbound
{
}

/** Request payload for Inbound#create. */
class InboundCreateData
{
    public ?string $token = null;
}

/** Markdown entity data model. */
class Markdown
{
    public ?string $filename = null;
    public string $markdown;
}

/** Request payload for Markdown#create. */
class MarkdownCreateData
{
    public ?string $filename = null;
    public string $markdown;
}

/** Men entity data model. */
class Men
{
}

/** Request payload for Men#load. */
class MenLoadMatch
{
}

/** Merge entity data model. */
class Merge
{
    public array $files;
}

/** Request payload for Merge#create. */
class MergeCreateData
{
    public array $files;
}

/** Portal entity data model. */
class Portal
{
}

/** Request payload for Portal#create. */
class PortalCreateData
{
}

/** RootGet entity data model. */
class RootGet
{
}

/** Request payload for RootGet#load. */
class RootGetLoadMatch
{
}

/** Signup entity data model. */
class Signup
{
    public string $email;
}

/** Request payload for Signup#create. */
class SignupCreateData
{
    public string $email;
}

/** Split entity data model. */
class Split
{
    public string $file;
}

/** Request payload for Split#create. */
class SplitCreateData
{
    public string $file;
}

/** Table entity data model. */
class Table
{
    public string $file;
}

/** Request payload for Table#create. */
class TableCreateData
{
    public string $file;
}

/** Template entity data model. */
class Template
{
    public array $data;
    public ?string $filename = null;
    public string $template;
}

/** Request payload for Template#create. */
class TemplateCreateData
{
    public array $data;
    public ?string $filename = null;
    public string $template;
}

/** Text entity data model. */
class Text
{
    public string $file;
}

/** Request payload for Text#create. */
class TextCreateData
{
    public string $file;
}

/** ToImage entity data model. */
class ToImage
{
    public ?int $dpi = null;
    public string $file;
}

/** Request payload for ToImage#create. */
class ToImageCreateData
{
    public ?int $dpi = null;
    public string $file;
}

/** Translate entity data model. */
class Translate
{
    public mixed $file = null;
    public ?string $output = null;
    public string $target_language;
    public mixed $text = null;
}

/** Request payload for Translate#create. */
class TranslateCreateData
{
    public mixed $file = null;
    public ?string $output = null;
    public string $target_language;
    public mixed $text = null;
}

/** Watermark entity data model. */
class Watermark
{
    public string $file;
    public ?string $text = null;
}

/** Request payload for Watermark#create. */
class WatermarkCreateData
{
    public string $file;
    public ?string $text = null;
}

/** Webhook entity data model. */
class Webhook
{
}

/** Request payload for Webhook#create. */
class WebhookCreateData
{
}

