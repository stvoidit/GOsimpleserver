module.exports = {
  devServer: {
    proxy: 'http://127.0.0.1:9000'
  },
  outputDir: '../../build/static/',
  assetsDir: '',
  pages: {
    login: {
      entry: 'src/Login/main.js',
      // filename: '../../build/static/templates/login.html',
      title: 'login',
      chunks: ['chunk-vendors', 'chunk-common', 'login']
    },
    MyVieos: {
      entry: 'src/MyVieos/main.js',
      // filename: '../../build/static/templates/MyVieos.html',
      title: 'MyVieos',
      chunks: ['chunk-vendors', 'chunk-common', 'MyVieos']
    }

  }
}