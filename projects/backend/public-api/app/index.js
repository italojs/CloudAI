'use strict'

const cors = require('cors')
const routes = require('./routes')
const appify = require('@mantris/appify')

/**
 * Application setup.
 * @param {Object} api                 Express instance.
 * @param {Object} options.config      Application configs.
 * @param {String} options.environment Environment name.
 */
module.exports = appify((api, config, environment) => {
  api.use(cors(config.cors))

  api.get('ms-crop-image/ping', routes.cropImage.ping.factory(config.msCropImage))
  api.post('ms-crop-image/files', routes.cropImage.create.factory(config.msCropImage))

  api.get('/cloud-storage/:file', routes.cloudStorage.find.factory(config.msCloudStorage))
  api.get('/cloud-storage/ping', routes.cloudStorage.ping.factory(config.msCloudStorage))
})
