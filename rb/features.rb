# InkspanDocument SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module InkspanDocumentFeatures
  def self.make_feature(name)
    case name
    when "base"
      InkspanDocumentBaseFeature.new
    when "ratelimit"
      InkspanDocumentRatelimitFeature.new
    when "retry"
      InkspanDocumentRetryFeature.new
    when "test"
      InkspanDocumentTestFeature.new
    when "timeout"
      InkspanDocumentTimeoutFeature.new
    else
      InkspanDocumentBaseFeature.new
    end
  end
end
