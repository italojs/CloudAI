'use strict'

const proxy = require('@mantris/proxy')

const factory = ({ url, timeout }) => {
  return proxy({
    to: url,
    at: '/create',
    timeout
  })
}

module.exports = { factory }
