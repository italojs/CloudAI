
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

describe('DELETE /provider-contracts/:unit/', () => {
  let api
  let endpoint

  before(() => {
    api = axiosist(app(config, 'testing'))

    endpoint = nock(config.provider.url)
      .delete('/contracts/5aa68aa49521341b7acc4347')
      .once()
      .reply()

    api.delete('/provider-contracts/5aa68aa49521341b7acc4347')
  })

  after(() => {
    nock.cleanAll()
  })

  it('redirects to /contracts/:unit', async () => {
    expect(endpoint.isDone()).to.be.equals(true)
  })
})
