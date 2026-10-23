# InkspanDocument SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

InkspanDocumentUtility.registrar = ->(u) {
  u.clean = InkspanDocumentUtilities::Clean
  u.clean_add = InkspanDocumentUtilities::CleanAdd
  u.clean_explain = InkspanDocumentUtilities::CleanExplain
  u.done = InkspanDocumentUtilities::Done
  u.make_error = InkspanDocumentUtilities::MakeError
  u.feature_add = InkspanDocumentUtilities::FeatureAdd
  u.feature_hook = InkspanDocumentUtilities::FeatureHook
  u.feature_init = InkspanDocumentUtilities::FeatureInit
  u.fetcher = InkspanDocumentUtilities::Fetcher
  u.make_fetch_def = InkspanDocumentUtilities::MakeFetchDef
  u.make_context = InkspanDocumentUtilities::MakeContext
  u.make_options = InkspanDocumentUtilities::MakeOptions
  u.make_request = InkspanDocumentUtilities::MakeRequest
  u.make_response = InkspanDocumentUtilities::MakeResponse
  u.make_result = InkspanDocumentUtilities::MakeResult
  u.make_point = InkspanDocumentUtilities::MakePoint
  u.make_spec = InkspanDocumentUtilities::MakeSpec
  u.make_url = InkspanDocumentUtilities::MakeUrl
  u.param = InkspanDocumentUtilities::Param
  u.prepare_auth = InkspanDocumentUtilities::PrepareAuth
  u.prepare_body = InkspanDocumentUtilities::PrepareBody
  u.prepare_headers = InkspanDocumentUtilities::PrepareHeaders
  u.prepare_method = InkspanDocumentUtilities::PrepareMethod
  u.prepare_params = InkspanDocumentUtilities::PrepareParams
  u.prepare_path = InkspanDocumentUtilities::PreparePath
  u.prepare_query = InkspanDocumentUtilities::PrepareQuery
  u.graphql_body = InkspanDocumentUtilities::GraphqlBody
  u.graphql_errors = InkspanDocumentUtilities::GraphqlErrors
  u.result_basic = InkspanDocumentUtilities::ResultBasic
  u.result_body = InkspanDocumentUtilities::ResultBody
  u.result_headers = InkspanDocumentUtilities::ResultHeaders
  u.transform_request = InkspanDocumentUtilities::TransformRequest
  u.transform_response = InkspanDocumentUtilities::TransformResponse
}
