# frozen_string_literal: true

# Typed models for the InkspanDocument SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Checkout entity data model.
class Checkout
end

# Request payload for Checkout#create.
#
# @!attribute [rw] plan
#   @return [String]
CheckoutCreateData = Struct.new(
  :plan,
  keyword_init: true
)

# Compress entity data model.
#
# @!attribute [rw] file
#   @return [String]
Compress = Struct.new(
  :file,
  keyword_init: true
)

# Request payload for Compress#create.
#
# @!attribute [rw] file
#   @return [String]
CompressCreateData = Struct.new(
  :file,
  keyword_init: true
)

# FillForm entity data model.
#
# @!attribute [rw] fields
#   @return [String]
#
# @!attribute [rw] file
#   @return [String]
FillForm = Struct.new(
  :fields,
  :file,
  keyword_init: true
)

# Request payload for FillForm#create.
#
# @!attribute [rw] fields
#   @return [String]
#
# @!attribute [rw] file
#   @return [String]
FillFormCreateData = Struct.new(
  :fields,
  :file,
  keyword_init: true
)

# Health entity data model.
class Health
end

# Request payload for Health#load.
class HealthLoadMatch
end

# Html entity data model.
#
# @!attribute [rw] filename
#   @return [String, nil]
#
# @!attribute [rw] html
#   @return [String]
Html = Struct.new(
  :filename,
  :html,
  keyword_init: true
)

# Request payload for Html#create.
#
# @!attribute [rw] filename
#   @return [String, nil]
#
# @!attribute [rw] html
#   @return [String]
HtmlCreateData = Struct.new(
  :filename,
  :html,
  keyword_init: true
)

# Inbound entity data model.
class Inbound
end

# Request payload for Inbound#create.
#
# @!attribute [rw] token
#   @return [String, nil]
InboundCreateData = Struct.new(
  :token,
  keyword_init: true
)

# Markdown entity data model.
#
# @!attribute [rw] filename
#   @return [String, nil]
#
# @!attribute [rw] markdown
#   @return [String]
Markdown = Struct.new(
  :filename,
  :markdown,
  keyword_init: true
)

# Request payload for Markdown#create.
#
# @!attribute [rw] filename
#   @return [String, nil]
#
# @!attribute [rw] markdown
#   @return [String]
MarkdownCreateData = Struct.new(
  :filename,
  :markdown,
  keyword_init: true
)

# Men entity data model.
class Men
end

# Request payload for Men#load.
class MenLoadMatch
end

# Merge entity data model.
#
# @!attribute [rw] files
#   @return [Array]
Merge = Struct.new(
  :files,
  keyword_init: true
)

# Request payload for Merge#create.
#
# @!attribute [rw] files
#   @return [Array]
MergeCreateData = Struct.new(
  :files,
  keyword_init: true
)

# Portal entity data model.
class Portal
end

# Request payload for Portal#create.
class PortalCreateData
end

# RootGet entity data model.
class RootGet
end

# Request payload for RootGet#load.
class RootGetLoadMatch
end

# Signup entity data model.
#
# @!attribute [rw] email
#   @return [String]
Signup = Struct.new(
  :email,
  keyword_init: true
)

# Request payload for Signup#create.
#
# @!attribute [rw] email
#   @return [String]
SignupCreateData = Struct.new(
  :email,
  keyword_init: true
)

# Split entity data model.
#
# @!attribute [rw] file
#   @return [String]
Split = Struct.new(
  :file,
  keyword_init: true
)

# Request payload for Split#create.
#
# @!attribute [rw] file
#   @return [String]
SplitCreateData = Struct.new(
  :file,
  keyword_init: true
)

# Table entity data model.
#
# @!attribute [rw] file
#   @return [String]
Table = Struct.new(
  :file,
  keyword_init: true
)

# Request payload for Table#create.
#
# @!attribute [rw] file
#   @return [String]
TableCreateData = Struct.new(
  :file,
  keyword_init: true
)

# Template entity data model.
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] filename
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [String]
Template = Struct.new(
  :data,
  :filename,
  :template,
  keyword_init: true
)

# Request payload for Template#create.
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] filename
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [String]
TemplateCreateData = Struct.new(
  :data,
  :filename,
  :template,
  keyword_init: true
)

# Text entity data model.
#
# @!attribute [rw] file
#   @return [String]
Text = Struct.new(
  :file,
  keyword_init: true
)

# Request payload for Text#create.
#
# @!attribute [rw] file
#   @return [String]
TextCreateData = Struct.new(
  :file,
  keyword_init: true
)

# ToImage entity data model.
#
# @!attribute [rw] dpi
#   @return [Integer, nil]
#
# @!attribute [rw] file
#   @return [String]
ToImage = Struct.new(
  :dpi,
  :file,
  keyword_init: true
)

# Request payload for ToImage#create.
#
# @!attribute [rw] dpi
#   @return [Integer, nil]
#
# @!attribute [rw] file
#   @return [String]
ToImageCreateData = Struct.new(
  :dpi,
  :file,
  keyword_init: true
)

# Translate entity data model.
#
# @!attribute [rw] file
#   @return [Object, nil]
#
# @!attribute [rw] output
#   @return [String, nil]
#
# @!attribute [rw] target_language
#   @return [String]
#
# @!attribute [rw] text
#   @return [Object, nil]
Translate = Struct.new(
  :file,
  :output,
  :target_language,
  :text,
  keyword_init: true
)

# Request payload for Translate#create.
#
# @!attribute [rw] file
#   @return [Object, nil]
#
# @!attribute [rw] output
#   @return [String, nil]
#
# @!attribute [rw] target_language
#   @return [String]
#
# @!attribute [rw] text
#   @return [Object, nil]
TranslateCreateData = Struct.new(
  :file,
  :output,
  :target_language,
  :text,
  keyword_init: true
)

# Watermark entity data model.
#
# @!attribute [rw] file
#   @return [String]
#
# @!attribute [rw] text
#   @return [String, nil]
Watermark = Struct.new(
  :file,
  :text,
  keyword_init: true
)

# Request payload for Watermark#create.
#
# @!attribute [rw] file
#   @return [String]
#
# @!attribute [rw] text
#   @return [String, nil]
WatermarkCreateData = Struct.new(
  :file,
  :text,
  keyword_init: true
)

# Webhook entity data model.
class Webhook
end

# Request payload for Webhook#create.
class WebhookCreateData
end

