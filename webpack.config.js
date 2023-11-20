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

// const path = require('path');

// module.exports = {
//   entry: './web/src/js/main.ts', // Путь к вашему TypeScript файлу
//   output: {
//     filename: 'bundle.js',
//     path: path.resolve(__dirname, 'dist'),
//   },
//   module: {
//     rules: [
//       {
//         test: /\.tsx?$/,
//         use: 'ts-loader',
//         exclude: /node_modules/,
//       },
//     ],
//   },
//   resolve: {
//     extensions: ['.tsx', '.ts', '.js'],
//   },
// };