module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  devServer: {
    proxy: {
      "/api": {
        target: 'http://localhost:8080'
      },
      "/.auth": {
        target: 'http://localhost:8080',
      }
    }
  }, 
  pages: {
    index: {
      entry: 'src/main.js',
      template: 'public/index.html',
      filename: 'index.html',
      title: 'Index Page',
      chunks: ['chunk-vendors', 'chunk-common', 'index']
    },
    login: 'src/login/login.js'
  }
}