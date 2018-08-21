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

describe('GET /client-companies', () => {
  let api
  let endpoint

  before(() => {
    api = axiosist(app(config, 'testing'))

    endpoint = nock(config.client.url)
      .get('/companies')
      .once()
      .reply()

    api.get('/client-companies')
  })

  after(() => {
    nock.cleanAll()
  })

  it('redirects to /companies', async () => {
    expect(endpoint.isDone()).to.be.equals(true)
  })
})
