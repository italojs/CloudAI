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

describe('POST /client-holdings', () => {
  let api
  let endpoint

  before(() => {
    api = axiosist(app(config, 'testing'))

    endpoint = nock(config.client.url)
      .post('/holdings')
      .once()
      .reply()

    api.post('/client-holdings')
  })

  after(() => {
    nock.cleanAll()
  })

  it('redirects to /holdings', async () => {
    expect(endpoint.isDone()).to.be.equals(true)
  })
})
