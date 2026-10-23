# RootGet entity test

require "minitest/autorun"
require "json"
require_relative "../InkspanDocument_sdk"
require_relative "runner"

class RootGetEntityTest < Minitest::Test
  # main.kit.test.live.strict is true (the default is true): a live
  # request that fails, or a live test missing an input it needs,
  # fails the test.
  # An account with no record for a test to read skips it either way.
  LIVE_STRICT = true

  def test_create_instance
    testsdk = InkspanDocumentSDK.test(nil, nil)
    ent = testsdk.RootGet(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = root_get_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "root_get." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    root_get_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.root_get")))
    root_get_ref01_data = nil
    if root_get_ref01_data_raw.length > 0
      root_get_ref01_data = Helpers.to_map(root_get_ref01_data_raw[0][1])
    end

    # LOAD
    root_get_ref01_ent = client.RootGet(nil)
    root_get_ref01_match_dt0 = {}
    root_get_ref01_data_dt0_loaded = root_get_ref01_ent.load(root_get_ref01_match_dt0, nil)
    assert !root_get_ref01_data_dt0_loaded.nil?

  end
end

def root_get_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "root_get", "RootGetTestData.json")
  entity_data_source = File.read(entity_data_file, encoding: "UTF-8")
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = InkspanDocumentSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["root_get01", "root_get02", "root_get03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Whether *_ENTID supplied the idmap, read before env_override consumes
  # it: without it, the ids a live flow binds are the fixture's synthetic ones.
  entid_env_raw = ENV["INKSPAN_DOCUMENT_TEST_ROOT_GET_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "INKSPAN_DOCUMENT_TEST_ROOT_GET_ENTID" => idmap,
    "INKSPAN_DOCUMENT_TEST_LIVE" => "FALSE",
    "INKSPAN_DOCUMENT_TEST_EXPLAIN" => "FALSE",
  })

  idmap_resolved = Helpers.to_map(
    env["INKSPAN_DOCUMENT_TEST_ROOT_GET_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["INKSPAN_DOCUMENT_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      # FIRST, so the generated fields below win: sdk-test-control.json's
      # test.client.options adds to the live client, it does not redirect it.
      Runner.live_client_options,
      {
      },
      extra || {},
    ])
    client = InkspanDocumentSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["INKSPAN_DOCUMENT_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["INKSPAN_DOCUMENT_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
