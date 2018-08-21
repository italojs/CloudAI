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

describe('GET /client-holdins/:holding', () => {
  let api
  let endpoint

  before(() => {
    api = axiosist(app(config, 'testing'))

    endpoint = nock(config.client.url)
      .get('/holdings/5adde33d1857023d1f42ead7')
      .once()
      .reply()

    api.get('/client-holdings/5adde33d1857023d1f42ead7')
  })

  after(() => {
    nock.cleanAll()
  })

  it('redirects to /holdings/:holding', async () => {
    expect(endpoint.isDone()).to.be.equals(true)
  })
})
