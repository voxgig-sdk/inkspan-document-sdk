# Typed models for the InkspanDocument SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Checkout(TypedDict):
    pass


class CheckoutCreateData(TypedDict):
    plan: str


class Compress(TypedDict):
    file: str


class CompressCreateData(TypedDict):
    file: str


class FillForm(TypedDict):
    fields: str
    file: str


class FillFormCreateData(TypedDict):
    fields: str
    file: str


class Health(TypedDict):
    pass


class HealthLoadMatch(TypedDict):
    pass


class HtmlRequired(TypedDict):
    html: str


class Html(HtmlRequired, total=False):
    filename: str


class HtmlCreateDataRequired(TypedDict):
    html: str


class HtmlCreateData(HtmlCreateDataRequired, total=False):
    filename: str


class Inbound(TypedDict):
    pass


class InboundCreateData(TypedDict, total=False):
    token: str


class MarkdownRequired(TypedDict):
    markdown: str


class Markdown(MarkdownRequired, total=False):
    filename: str


class MarkdownCreateDataRequired(TypedDict):
    markdown: str


class MarkdownCreateData(MarkdownCreateDataRequired, total=False):
    filename: str


class Men(TypedDict):
    pass


class MenLoadMatch(TypedDict):
    pass


class Merge(TypedDict):
    files: list


class MergeCreateData(TypedDict):
    files: list


class Portal(TypedDict):
    pass


class PortalCreateData(TypedDict):
    pass


class RootGet(TypedDict):
    pass


class RootGetLoadMatch(TypedDict):
    pass


class Signup(TypedDict):
    email: str


class SignupCreateData(TypedDict):
    email: str


class Split(TypedDict):
    file: str


class SplitCreateData(TypedDict):
    file: str


class Table(TypedDict):
    file: str


class TableCreateData(TypedDict):
    file: str


class TemplateRequired(TypedDict):
    data: dict
    template: str


class Template(TemplateRequired, total=False):
    filename: str


class TemplateCreateDataRequired(TypedDict):
    data: dict
    template: str


class TemplateCreateData(TemplateCreateDataRequired, total=False):
    filename: str


class Text(TypedDict):
    file: str


class TextCreateData(TypedDict):
    file: str


class ToImageRequired(TypedDict):
    file: str


class ToImage(ToImageRequired, total=False):
    dpi: int


class ToImageCreateDataRequired(TypedDict):
    file: str


class ToImageCreateData(ToImageCreateDataRequired, total=False):
    dpi: int


class TranslateRequired(TypedDict):
    target_language: str


class Translate(TranslateRequired, total=False):
    file: Any
    output: str
    text: Any


class TranslateCreateDataRequired(TypedDict):
    target_language: str


class TranslateCreateData(TranslateCreateDataRequired, total=False):
    file: Any
    output: str
    text: Any


class WatermarkRequired(TypedDict):
    file: str


class Watermark(WatermarkRequired, total=False):
    text: str


class WatermarkCreateDataRequired(TypedDict):
    file: str


class WatermarkCreateData(WatermarkCreateDataRequired, total=False):
    text: str


class Webhook(TypedDict):
    pass


class WebhookCreateData(TypedDict):
    pass
