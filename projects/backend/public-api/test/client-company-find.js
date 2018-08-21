'use strict'

const nock = require('nock')
const app = require('../app')
const { expect } = require('chai')
const axiosist = require('axiosist')
const merge = require('lodash.merge')

const config = merge(
  require('../config'),
  { client: { url: 'http://srv-client.mock' } }
)

describe('GET /client-companies/:company', () => {
  let api
  let endpoint

  before(() => {
    api = axiosist(app(config, 'testing'))

    endpoint = nock(config.client.url)
      .get('/companies/5ac4fae258f6f2238118dc12')
      .once()
      .reply()

    api.get('/client-companies/5ac4fae258f6f2238118dc12')
  })

  after(() => {
    nock.cleanAll()
  })

  it('redirects to /companies/:company', async () => {
    expect(endpoint.isDone()).to.be.equals(true)
  })
})
