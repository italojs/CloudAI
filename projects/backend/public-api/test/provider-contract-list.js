/* global describe before after it */
'use strict'

const nock = require('nock')
const app = require('../app')
const { expect } = require('chai')
const axiosist = require('axiosist')
const merge = require('lodash.merge')

const config = merge(
  require('../config'),
  { provider: { url: 'http://api-private.mock' } }
)

describe('GET /provider-contracts/:unit/revisions', () => {
  let api
  let endpoint

  before(() => {
    api = axiosist(app(config, 'testing'))

    endpoint = nock(config.provider.url)
      .get('/contracts/5aa163072d8690238e5f458c/revisions')
      .once()
      .reply()

    api.get('/provider-contracts/5aa163072d8690238e5f458c/revisions')
  })

  after(() => {
    nock.cleanAll()
  })

  it('redirects to /contracts/:unit/revisions', async () => {
    expect(endpoint.isDone()).to.be.equals(true)
  })
})
