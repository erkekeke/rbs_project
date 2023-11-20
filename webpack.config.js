// webpack.config.js
const path = require('path');

module.exports = {
  entry: {
    main: './web/src/js/main.js'
    // Добавьте другие точки входа, если необходимо
  },
  output: {
    filename: '[name].bundle.js',
    path: path.resolve(__dirname, 'dist'),
  },
};