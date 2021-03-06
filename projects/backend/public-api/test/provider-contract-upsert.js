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

describe('PUT /provider-contracts/:unit/', () => {
  let api
  let endpoint

  before(() => {
    api = axiosist(app(config, 'testing'))

    endpoint = nock(config.provider.url)
      .put('/contracts/5aa163072d8690238e5f458c')
      .once()
      .reply()

    api.put('/provider-contracts/5aa163072d8690238e5f458c')
  })

  after(() => {
    nock.cleanAll()
  })

  it('redirects to /contracts/:unit', async () => {
    expect(endpoint.isDone()).to.be.equals(true)
  })
})
