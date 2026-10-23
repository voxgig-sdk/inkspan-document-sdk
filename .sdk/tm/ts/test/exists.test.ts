
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { InkspanDocumentSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = InkspanDocumentSDK.test()
    equal(testsdk instanceof InkspanDocumentSDK, true,
      'InkspanDocumentSDK.test() must return a client synchronously')
  })

})
