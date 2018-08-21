'use strict'

const env = require('sugar-env')

module.exports = {
  auth: {
    jwt: {
      issuer: env.get('AUTHORITY_ISSUER_URN'),
      audience: env.get('AUTHORITY_AUDIENCE_URN')
    },
    jwks: {
      uri: env.get('AUTH_JWKS_URI')
    }
  },
  cors: {
    origin: '*',
    methods: [ 'DELETE', 'GET', 'POST', 'PUT' ],
    preflightContinue: false,
    optionsSuccessStatus: 204
  },
  msCropImage: {
    url: env.get('CROP_IMAGE_URL'),
    timeout: parseInt(env.get('CROP_IMAGE_TIMEOUT', 3000))
  },
  multer: {
    fileSize: parseInt(env.get('MULTER_FILE_SIZE', 5 * 1024 * 1024))
  },
  msCloudStorage:{
    url: env.get('CLOUD_STORAGE_URL'),
    timeout: parseInt(env.get('CLOUD_STORAGE_TIMEOUT', 3000))
  }
}
