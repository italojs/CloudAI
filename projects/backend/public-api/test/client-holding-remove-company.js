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

const holding = '5adde33d1857023d1f42ead7'
const company = '5adf2bc606587001e202682a'

describe('DELETE /client-holdings/:holding/companies/:company', () => {
  let api
  let endpoint

  before(() => {
    api = axiosist(app(config, 'testing'))

    endpoint = nock(config.client.url)
      .delete(`/holdings/${holding}/companies/${company}`)
      .once()
      .reply()

    api.delete(`/client-holdings/${holding}/companies/${company}`)
  })

  after(() => {
    nock.cleanAll()
  })

  it('redirects to /holdings/:holding/companies/:company', async () => {
    expect(endpoint.isDone()).to.be.equals(true)
  })
})
