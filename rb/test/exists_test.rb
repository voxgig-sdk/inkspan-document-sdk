# InkspanDocument SDK exists test

require "minitest/autorun"
require_relative "../InkspanDocument_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = InkspanDocumentSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
