module.exports = {
  development: {
    isProduction: false,
    port: process.env.FRONT_PORT || 8080,
    apiPort: process.env.API_PORT || 8100,
    apiPath: '/v1',
    app: {
      name: 'Grox (development)'
    }
  },
  production: {
    isProduction: true,
    port: process.env.FRONT_PORT || 8080,
    apiPort: process.env.API_PORT || 8100,
    apiPath: '/v1',
    app: {
      name: 'Grox (production)'
    }
  }
}[process.env.NODE_ENV || 'development'];
