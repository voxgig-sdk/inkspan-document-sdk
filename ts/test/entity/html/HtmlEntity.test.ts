

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { InkspanDocumentSDK, BaseFeature, config, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('HtmlEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when INKSPAN_DOCUMENT_TEST_LIVE=TRUE.
  afterEach(liveDelay('INKSPAN_DOCUMENT_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = InkspanDocumentSDK.test()
    const ent = testsdk.Html()
    assert(null != ent)
  })


  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = InkspanDocumentSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.Html().create({"filename":1,"html":"x"} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.INKSPAN_DOCUMENT_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'html.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"filename":{"a":true,"h":"Filename","n":"filename","r":false,"t":"`$STRING`","key$":"filename","index$":0},"html":{"a":true,"h":"Html","n":"html","r":true,"t":"`$STRING`","key$":"html","index$":1}},"name":"html","op":{"create":{"input":"data","name":"create","points":[{"a":true,"bf":["filename","html"],"co":{"id":"POST /v1/generate/html","source":"openapi3","version":2},"g":{"header":[{"a":true,"k":"header","n":"x_api_key","or":"x-api-key","r":false,"t":"`$ANY`","index$":0},{"a":true,"k":"header","n":"x_rapidapi_proxy_secret","or":"x-rapidapi-proxy-secret","r":false,"t":"`$ANY`","index$":1}]},"k":"http","m":"POST","o":"/v1/generate/html","q":{},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"v1"},{"lit":"generate"},{"lit":"html"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"html","name__orig":"html","Name":"Html","name_":"html","name-":"html","NAME":"HTML","index$":4}, {"active":true,"entity":"html","key$":"BasicHtmlFlow","kind":"basic","name":"BasicHtmlFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"html_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'Html', {"POST /v1/generate/html":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"properties":{"html":{"type":"string","title":"Html","key$":"html"},"filename":{"type":"string","title":"Filename","default":"document.pdf","key$":"filename"}},"type":"object","required":["html"],"title":"HtmlReq","x-ref":"#/components/schemas/HtmlReq","index$":1}}}},"parameters":[{"name":"x-api-key","in":"header","required":false,"schema":{"anyOf":[{"type":"string"},{"type":"null"}],"title":"X-Api-Key"},"index$":0},{"name":"x-rapidapi-proxy-secret","in":"header","required":false,"schema":{"anyOf":[{"type":"string"},{"type":"null"}],"title":"X-Rapidapi-Proxy-Secret"},"index$":1}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const html_ref01_ent = client.Html()
    let html_ref01_data = setup.data.new.html['html_ref01']

    html_ref01_data = (await html_ref01_ent.create(html_ref01_data)).data()
    assert(null != html_ref01_data)


  })
})



// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const LIVE_STRICT = true

function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/html/HtmlTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = InkspanDocumentSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['html01','html02','html03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'INKSPAN_DOCUMENT_TEST_HTML_ENTID': idmap,
    'INKSPAN_DOCUMENT_TEST_LIVE': 'FALSE',
    'INKSPAN_DOCUMENT_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['INKSPAN_DOCUMENT_TEST_HTML_ENTID']

  const live = 'TRUE' === env.INKSPAN_DOCUMENT_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['INKSPAN_DOCUMENT_TEST_HTML_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new InkspanDocumentSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.INKSPAN_DOCUMENT_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
