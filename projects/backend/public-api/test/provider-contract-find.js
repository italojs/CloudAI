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

describe('GET /provider-contracts/:unit/revisions/:revision', () => {
  let api
  let endpoint

  before(() => {
    api = axiosist(app(config, 'testing'))

    endpoint = nock(config.provider.url)
      .get('/contracts/5aa68aa49521341b7acc4347/revisions/5aa68af79521341b7acc4348')
      .once()
      .reply()

    api.get('/provider-contracts/5aa68aa49521341b7acc4347/revisions/5aa68af79521341b7acc4348')
  })

  after(() => {
    nock.cleanAll()
  })

  it('redirects to /contracts/:unit/revisions/:revision', async () => {
    expect(endpoint.isDone()).to.be.equals(true)
  })
})
