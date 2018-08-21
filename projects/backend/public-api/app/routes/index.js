'use strict'

module.exports = {
  cloudStorage: {
    find: require('./cloud-storage/find'),
    ping: require('./cloud-storage/ping')
  },
  cropImage: {
    create: require('./crop-image/create'),
    ping: require('./crop-image/ping')
  }
}
