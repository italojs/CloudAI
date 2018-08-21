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

describe('DELETE /client-holdings/:holding', () => {
  let api
  let endpoint

  before(() => {
    api = axiosist(app(config, 'testing'))

    endpoint = nock(config.client.url)
      .delete('/holdings/5adde33d1857023d1f42ead7')
      .once()
      .reply()

    api.delete('/client-holdings/5adde33d1857023d1f42ead7')
  })

  after(() => {
    nock.cleanAll()
  })

  it('redirects to /holdings/:holding', async () => {
    expect(endpoint.isDone()).to.be.equals(true)
  })
})
