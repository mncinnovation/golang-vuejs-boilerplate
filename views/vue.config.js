const HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
    publicPath: '/',
    configureWebpack: {
        plugins: [
            new HtmlWebpackPlugin({
              title: 'Custom template',
              baseUrl: '/',
              // Load a custom template (lodash by default)
              template: 'public/index.html'
            })
        ]
    }
}  