# InkspanDocument SDK utility: make_context
require_relative '../core/context'
module InkspanDocumentUtilities
  MakeContext = ->(ctxmap, basectx) {
    InkspanDocumentContext.new(ctxmap, basectx)
  }
end
