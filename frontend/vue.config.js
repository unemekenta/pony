module.exports = {
  runtimeCompiler: true,
  devServer: {
    proxy: {
      '^/api/': {
          target: 'http://backend:8000',
          logLevel: 'debug',
          pathRewrite: { '^/api/': '/api/' }
      }
    }
  }
};