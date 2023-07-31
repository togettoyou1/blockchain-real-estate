//开发环境配置
'use strict'
// 配置文件合并模块
const merge = require('webpack-merge')
// 导入 prod.env.js 配置文件
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  // 导出一个对象，NODE_ENV 是一个环境变量，指定 development 环境
  NODE_ENV: '"development"'
})
